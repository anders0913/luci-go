// Copyright 2015 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package frontend

import (
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"github.com/julienschmidt/httprouter"
	"github.com/luci/luci-go/server/templates"
	"golang.org/x/net/context"

	"github.com/luci/luci-go/appengine/cmd/milo/settings"
	"github.com/luci/luci-go/appengine/cmd/milo/swarming"
	"github.com/luci/luci-go/appengine/gaeauth/server"
)

// Where it all begins!!!
func init() {
	// Register endpoint services.
	ss := &swarming.Service{}
	api, err := endpoints.RegisterService(ss, "swarming", "v1", "Milo Swarming API", true)
	if err != nil {
		log.Printf("Unable to register endpoint services: %s", err)
	} else {
		register := func(orig, name, method, path, desc string) {
			m := api.MethodByName(orig)
			i := m.Info()
			i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
		}
		register("Build", "swarming.build", "GET", "swarming", "Swarming Build view.")
	}

	// Register plain ol' http services.
	r := httprouter.New()
	server.InstallHandlers(r, settings.Base)
	r.GET("/", wrap(dummy{}))
	r.GET("/swarming/:server/:id/steps/*logname", wrap(swarming.Log{}))
	r.GET("/swarming/:server/:id", wrap(swarming.Build{}))

	// User settings
	r.GET("/settings", wrap(settings.Settings{}))
	r.POST("/settings", wrap(settings.Settings{}))

	http.Handle("/", r)

	endpoints.HandleHTTP()
}

type dummy struct{}

func (d dummy) GetTemplateName(t settings.Theme) string {
	return "base.html"
}

func (d dummy) Render(c context.Context, r *http.Request, p httprouter.Params) (*templates.Args, error) {
	return &templates.Args{"contents": "This is the root page"}, nil
}

// Do all the middleware initilization and theme handling.
func wrap(h settings.ThemedHandler) httprouter.Handle {
	return settings.Wrap(h)
}
