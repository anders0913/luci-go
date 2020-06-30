// Copyright 2020 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package partition encapsulates partitioning and querying large keyspace which
// can't be expressed even as uint64.
//
// All to/from string functions use hex encoding.
package partition

import (
	"fmt"
	"math/big"
	"strings"

	"go.chromium.org/luci/common/errors"
)

// Partition represents a range [Low..High).
type Partition struct {
	Low  big.Int // inclusive
	High big.Int // exclusive. May be equal to max SHA2 hash value + 1.
}

// SortedPartitions are disjoint partitions sorted by ascending .Low field.
type SortedPartitions []*Partition

func FromInts(low, high int64) *Partition {
	if low > high {
		panic(errors.Reason("Partition %d..%d is invalid", low, high))
	}
	p := &Partition{}
	p.Low.SetInt64(low)
	p.High.SetInt64(high)
	return p
}

func SpanInclusive(low, highInclusive string) (*Partition, error) {
	p := &Partition{}
	if err := setBigIntFromString(&p.Low, low); err != nil {
		return nil, err
	}
	if err := setBigIntFromString(&p.High, highInclusive); err != nil {
		return nil, err
	}
	p.High.Add(&p.High, bigInt1) // s.high++
	if p.Low.Cmp(&p.High) > 0 {
		return nil, errors.Reason("Partition %s is invalid", p.String()).Err()
	}
	return p, nil
}

func Universe(keySpaceBytes int) *Partition {
	p := &Partition{}
	p.High.SetBit(&p.High, keySpaceBytes*8, 1) // 2^(keySpaceBytes*8)
	return p
}

func FromString(s string) (*Partition, error) {
	i := strings.Index(s, "_")
	if i <= 0 || i == len(s)-1 {
		return nil, errors.Reason("partition %q has invalid format", s).Err()
	}
	p := &Partition{}
	if err := setBigIntFromString(&p.Low, s[:i]); err != nil {
		return nil, err
	}
	if err := setBigIntFromString(&p.High, s[i+1:]); err != nil {
		return nil, err
	}
	if p.Low.Cmp(&p.High) > 0 {
		return nil, errors.Reason("Partition %s is invalid", p.String()).Err()
	}
	return p, nil
}

func (p Partition) String() string {
	return fmt.Sprintf("%s_%s", p.Low.Text(16 /*hex*/), p.High.Text(16 /*hex*/))
}

func (p Partition) Copy() *Partition {
	r := &Partition{}
	r.Low.Set(&p.Low)
	r.High.Set(&p.High)
	return r
}

func (p Partition) QueryBounds(keySpaceBytes int) (low, high string) {
	low = paddedHex(&p.Low, keySpaceBytes)
	if inKeySpace(&p.High, keySpaceBytes) {
		// In practice, this should mean p.high == 2^(keySpaceBytes*8).
		high = "g" // all hex strings are smaller than "g".
	} else {
		high = paddedHex(&p.High, keySpaceBytes)
	}
	return
}

func (p Partition) Split(shards int) SortedPartitions {
	if shards <= 0 {
		panic(">=1 shard required")
	}
	var increment, remainder, cur big.Int
	increment.QuoRem(
		cur.Sub(&p.High, &p.Low),
		big.NewInt(int64(shards)),
		&remainder)
	if remainder.Cmp(bigInt0) > 0 {
		increment.Add(&increment, bigInt1)
	}

	partitions := make([]*Partition, 0, shards)
	cur.Set(&p.Low)
	for cur.Cmp(&p.High) < 0 {
		next := &Partition{}
		next.Low.Set(&cur)
		next.High.Add(&cur, &increment)
		cur.Set(&next.High)
		partitions = append(partitions, next)
	}
	// Due to int division to compute the increment, the last partition may
	// overshoot, so ensure it ends exactly at the end of the original.
	partitions[len(partitions)-1].High = p.High
	return partitions
}

// EducatedSplitAfter splits partition after a given boundary assuming constant
// density s.t. each shard has approximately targetItems.
//
// Caps the number of resulting partitions to at most maxShards.
// panics if called on invalid data.
func (p Partition) EducatedSplitAfter(exclusive string, beforeItems, targetItems, maxShards int) SortedPartitions {
	remaining := Partition{}
	if err := setBigIntFromString(&remaining.Low, exclusive); err != nil {
		panic(err)
	}
	if p.Low.Cmp(&remaining.Low) > 0 { // low > remaining.Low
		panic("must be within the partition")
	}
	if p.High.Cmp(&remaining.Low) <= 0 { // high <= remaining.Low
		panic("must be within the partition")
	}
	remaining.Low.Add(&remaining.Low, bigInt1) // remaining.Low++
	remaining.High.Set(&p.High)

	// Compute expShards as
	//
	//     beforeItems / len(before) * len(remaining) / targetItems
	//
	// in a somewhat readable way as
	//
	//     (beforeItems * len(remaining)) / ( targetItems * len(before))
	//
	// NOTE: this can be optimized if needed to avoid excessive memory allocations
	// in bit.Int at the cost of readability.
	iBefore := big.NewInt(int64(beforeItems))
	iTarget := big.NewInt(int64(targetItems))
	var expShards, iRemainder big.Int
	expShards.QuoRem(
		(&big.Int{}).Mul(iBefore, distance(&remaining.Low, &remaining.High)),
		(&big.Int{}).Mul(iTarget, distance(&p.Low, &remaining.Low)),
		&iRemainder,
	)
	if iRemainder.Cmp(bigInt0) > 0 {
		expShards.Add(&expShards, bigInt1)
	}
	shards := maxShards
	if expShards.Cmp(big.NewInt(int64(maxShards))) < 0 {
		shards = int(expShards.Int64())
	}
	return remaining.Split(shards)
}

var (
	// these are effectively constants predefined to avoid needless memory allocations.

	bigInt0 = big.NewInt(0)
	bigInt1 = big.NewInt(1)
)

func distance(low, high *big.Int) *big.Int {
	return (&big.Int{}).Sub(high, low)
}

func setBigIntFromString(b *big.Int, s string) error {
	if _, ok := b.SetString(s, 16 /*hex*/); !ok {
		return errors.Reason("invalid bigint hex %q", s).Err()
	}
	if b.Sign() == -1 {
		return errors.Reason("negative value %q not allowed", s).Err()
	}
	return nil
}

func paddedHex(b *big.Int, keySpaceBytes int) string {
	s := b.Text(16 /*hex*/)
	return strings.Repeat("0", keySpaceBytes*2-len(s)) + s
}

// inKeySpace returns whether v does not exceed keyspace upper boundary.
func inKeySpace(v *big.Int, keySpaceBytes int) bool {
	return v.BitLen() >= keySpaceBytes*8
}
