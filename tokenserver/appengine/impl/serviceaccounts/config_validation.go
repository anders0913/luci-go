// Copyright 2017 The LUCI Authors.
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

package serviceaccounts

import (
	"fmt"
	"strings"

	"github.com/luci/luci-go/common/config/validation"
	"github.com/luci/luci-go/common/data/stringset"
	"github.com/luci/luci-go/server/auth/identity"

	"github.com/luci/luci-go/tokenserver/api/admin/v1"
	"github.com/luci/luci-go/tokenserver/appengine/impl/utils/policy"
)

// validateConfigs validates the structure of configs fetched by fetchConfigs.
func validateConfigs(bundle policy.ConfigBundle, ctx *validation.Context) {
	ctx.SetFile(serviceAccountsCfg)
	cfg, ok := bundle[serviceAccountsCfg].(*admin.ServiceAccountsPermissions)
	if !ok {
		ctx.Error("unexpectedly wrong proto type %T", cfg)
		return
	}

	names := stringset.New(0)
	accounts := map[string]string{} // service account -> rule name where its defined
	for i, rule := range cfg.Rules {
		// Rule name must be unique. Missing name will be handled by 'validateRule'.
		if rule.Name != "" {
			if names.Has(rule.Name) {
				ctx.Error("two rules with identical name %q", rule.Name)
			} else {
				names.Add(rule.Name)
			}
		}

		// There should be no overlap between service account sets covered by each
		// rule.
		for _, account := range rule.ServiceAccount {
			if name, ok := accounts[account]; ok {
				ctx.Error("service account %q is mentioned by more than one rule (%q and %q)", account, name, rule.Name)
			} else {
				accounts[account] = rule.Name
			}
		}

		validateRule(fmt.Sprintf("rule #%d: %q", i+1, rule.Name), rule, ctx)
	}
}

// validateRule checks single ServiceAccountRule proto.
func validateRule(title string, r *admin.ServiceAccountRule, ctx *validation.Context) {
	ctx.Enter(title)
	defer ctx.Exit()

	if r.Name == "" {
		ctx.Error(`"name" is required`)
	}

	// Note: we allow any of the sets to be empty. The rule will just not match
	// anything in this case, this is fine.
	validateEmails("service_account", r.ServiceAccount, ctx)
	validateScopes("allowed_scope", r.AllowedScope, ctx)
	validateIdSet("end_user", r.EndUser, ctx)
	validateIdSet("proxy", r.Proxy, ctx)

	if r.MaxGrantValidityDuration != 0 {
		switch {
		case r.MaxGrantValidityDuration < 0:
			ctx.Error(`"max_grant_validity_duration" must be positive`)
		case r.MaxGrantValidityDuration > maxAllowedMaxGrantValidityDuration:
			ctx.Error(`"max_grant_validity_duration" must not exceed %d`, maxAllowedMaxGrantValidityDuration)
		}
	}
}

func validateEmails(field string, emails []string, ctx *validation.Context) {
	ctx.Enter("%q", field)
	defer ctx.Exit()
	for _, email := range emails {
		// We reuse 'user:' identity validator, user identities are emails too.
		if _, err := identity.MakeIdentity("user:" + email); err != nil {
			ctx.Error("bad email %q - %s", email, err)
		}
	}
}

func validateScopes(field string, scopes []string, ctx *validation.Context) {
	ctx.Enter("%q", field)
	defer ctx.Exit()
	for _, scope := range scopes {
		if !strings.HasPrefix(scope, "https://www.googleapis.com/") {
			ctx.Error("bad scope %q", scope)
		}
	}
}

func validateIdSet(field string, ids []string, ctx *validation.Context) {
	ctx.Enter("%q", field)
	defer ctx.Exit()
	for _, entry := range ids {
		if strings.HasPrefix(entry, "group:") {
			if entry[len("group:"):] == "" {
				ctx.Error("bad group entry - no group name")
			}
		} else if _, err := identity.MakeIdentity(entry); err != nil {
			ctx.Error("bad identity %q - %s", entry, err)
		}
	}
}
