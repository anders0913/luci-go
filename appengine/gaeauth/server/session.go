// Copyright 2015 The LUCI Authors.
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

package server

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	ds "go.chromium.org/gae/service/datastore"
	"go.chromium.org/gae/service/info"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/server/auth"
)

// SessionStore stores auth sessions in the datastore (always in the default
// namespace). It implements auth.SessionStore.
type SessionStore struct {
	Prefix string // used as prefix for datastore keys
}

// defaultNS returns GAE context configured to use default namespace.
func defaultNS(ctx context.Context) context.Context {
	return info.MustNamespace(ctx, "")
}

// OpenSession create a new session for a user with given expiration time.
// It returns unique session ID.
func (s *SessionStore) OpenSession(ctx context.Context, userID string, u *auth.User, exp time.Time) (string, error) {
	if strings.IndexByte(s.Prefix, '/') != -1 {
		return "", fmt.Errorf("gaeauth: bad prefix (%q) in SessionStore", s.Prefix)
	}
	if strings.IndexByte(userID, '/') != -1 {
		return "", fmt.Errorf("gaeauth: bad userID (%q), cannot have '/' inside", userID)
	}
	if err := u.Identity.Validate(); err != nil {
		return "", fmt.Errorf("gaeauth: bad identity string (%q) - %s", u.Identity, err)
	}
	ctx = defaultNS(ctx)

	now := clock.Now(ctx).UTC()
	prof := profile{
		Identity:  string(u.Identity),
		Superuser: u.Superuser,
		Email:     u.Email,
		Name:      u.Name,
		Picture:   u.Picture,
	}

	// Set in the transaction below.
	var sessionID string

	err := ds.RunInTransaction(ds.WithoutTransaction(ctx), func(ctx context.Context) error {
		// Grab existing userEntity or initialize a new one.
		userEnt := userEntity{ID: s.Prefix + "/" + userID}
		err := ds.Get(ctx, &userEnt)
		if err != nil && err != ds.ErrNoSuchEntity {
			return err
		}
		if err == ds.ErrNoSuchEntity {
			userEnt.Profile = prof
			userEnt.Created = now
		}
		userEnt.LastLogin = now

		// Make new session. ID will be generated by the datastore.
		sessionEnt := sessionEntity{
			Parent:     ds.KeyForObj(ctx, &userEnt),
			Profile:    prof,
			Created:    now,
			Expiration: exp.UTC(),
		}
		if err = ds.Put(ctx, &userEnt, &sessionEnt); err != nil {
			return err
		}

		sessionID = fmt.Sprintf("%s/%s/%d", s.Prefix, userID, sessionEnt.ID)
		return nil
	}, nil)

	if err != nil {
		return "", transient.Tag.Apply(err)
	}
	return sessionID, nil
}

// CloseSession closes a session given its ID. Does nothing if session is
// already closed or doesn't exist. Returns only transient errors.
func (s *SessionStore) CloseSession(ctx context.Context, sessionID string) error {
	ctx = defaultNS(ctx)
	ent, err := s.fetchSession(ctx, sessionID)
	switch {
	case err != nil:
		return err
	case ent == nil:
		return nil
	default:
		ent.IsClosed = true
		ent.Closed = clock.Now(ctx).UTC()
		return transient.Tag.Apply(ds.Put(ds.WithoutTransaction(ctx), ent))
	}
}

// GetSession returns existing non-expired session given its ID. Returns nil
// if session doesn't exist, closed or expired. Returns only transient errors.
func (s *SessionStore) GetSession(ctx context.Context, sessionID string) (*auth.Session, error) {
	ctx = defaultNS(ctx)
	ent, err := s.fetchSession(ctx, sessionID)
	if ent == nil {
		return nil, err
	}
	return &auth.Session{
		SessionID: sessionID,
		UserID:    ent.Parent.StringID()[len(s.Prefix)+1:],
		User: auth.User{
			Identity:  identity.Identity(ent.Profile.Identity),
			Superuser: ent.Profile.Superuser,
			Email:     ent.Profile.Email,
			Name:      ent.Profile.Name,
			Picture:   ent.Profile.Picture,
		},
		Exp: ent.Expiration,
	}, nil
}

// fetchSession fetches sessionEntity from the datastore and returns it if it is
// still open and non-expired. Returns (nil, nil) otherwise. Broken sessionID is
// logged and ignored, the function returns (nil, nil) in such case. Returns
// only transient errors.
func (s *SessionStore) fetchSession(ctx context.Context, sessionID string) (*sessionEntity, error) {
	chunks := strings.Split(sessionID, "/")
	if len(chunks) != 3 || chunks[0] != s.Prefix {
		logging.Warningf(ctx, "Malformed session ID %q, ignoring", sessionID)
		return nil, nil
	}
	id, err := strconv.ParseInt(chunks[2], 10, 64)
	if err != nil {
		logging.Warningf(ctx, "Malformed session ID %q, ignoring", sessionID)
		return nil, nil
	}

	ctx = ds.WithoutTransaction(ctx)
	sessionEnt := sessionEntity{
		ID:     id,
		Parent: ds.MakeKey(ctx, "gaeauth.User", chunks[0]+"/"+chunks[1]),
	}
	switch err = ds.Get(ctx, &sessionEnt); err {
	case nil:
		if sessionEnt.IsClosed || clock.Now(ctx).After(sessionEnt.Expiration) {
			return nil, nil
		}
		return &sessionEnt, nil
	case ds.ErrNoSuchEntity:
		return nil, nil
	default:
		return nil, transient.Tag.Apply(err)
	}
}

////

// profile is used in both userEntity and sessionEntity. It holds information
// about a user extracted from user.User struct.
type profile struct {
	Identity  string
	Superuser bool   `gae:",noindex"`
	Email     string `gae:",noindex"`
	Name      string `gae:",noindex"`
	Picture   string `gae:",noindex"`
}

// userEntity holds profile information of some user. It is root entity.
// ID is "<prefix>/<userID>" where <prefix> is SessionStore.Prefix, and <userID>
// is what is passed to OpenSession (unique user id as returned by
// authentication backend). Created or refreshed in OpenSession.
type userEntity struct {
	_kind string `gae:"$kind,gaeauth.User"`

	ID string `gae:"$id"`

	Profile   profile
	Created   time.Time // when this entity was created
	LastLogin time.Time // when last session was opened
}

// sessionEntity stores session information associated with session cookie.
// Parent entity is userEntity, ID is generated by the datastore. Includes user
// profile info inline to avoid additional datastore calls in GetSession. Never
// deleted from the datastore (to keep some sort of history of logins). Marked
// as closed in CloseSession. Since all user's sessions belong to single entity
// group, there's implicit 1 login per second per user limit on rate of logins.
type sessionEntity struct {
	_kind string `gae:"$kind,gaeauth.Session"`

	ID     int64   `gae:"$id"`
	Parent *ds.Key `gae:"$parent"`

	Profile    profile
	Created    time.Time // when this session was created
	Expiration time.Time // when this session expires

	IsClosed bool      // if true, the session was closed by CloseSession()
	Closed   time.Time // when the session was closed by CloseSession()
}
