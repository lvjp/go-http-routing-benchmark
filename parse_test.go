// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

// Parse
// https://parse.com/docs/rest#summary
var parseAPI = []router.Route{
	// Objects
	{Method: "POST", Path: "/1/classes/:className"},
	{Method: "GET", Path: "/1/classes/:className/:objectId"},
	{Method: "PUT", Path: "/1/classes/:className/:objectId"},
	{Method: "GET", Path: "/1/classes/:className"},
	{Method: "DELETE", Path: "/1/classes/:className/:objectId"},

	// Users
	{Method: "POST", Path: "/1/users"},
	{Method: "GET", Path: "/1/login"},
	{Method: "GET", Path: "/1/users/:objectId"},
	{Method: "PUT", Path: "/1/users/:objectId"},
	{Method: "GET", Path: "/1/users"},
	{Method: "DELETE", Path: "/1/users/:objectId"},
	{Method: "POST", Path: "/1/requestPasswordReset"},

	// Roles
	{Method: "POST", Path: "/1/roles"},
	{Method: "GET", Path: "/1/roles/:objectId"},
	{Method: "PUT", Path: "/1/roles/:objectId"},
	{Method: "GET", Path: "/1/roles"},
	{Method: "DELETE", Path: "/1/roles/:objectId"},

	// Files
	{Method: "POST", Path: "/1/files/:fileName"},

	// Analytics
	{Method: "POST", Path: "/1/events/:eventName"},

	// Push Notifications
	{Method: "POST", Path: "/1/push"},

	// Installations
	{Method: "POST", Path: "/1/installations"},
	{Method: "GET", Path: "/1/installations/:objectId"},
	{Method: "PUT", Path: "/1/installations/:objectId"},
	{Method: "GET", Path: "/1/installations"},
	{Method: "DELETE", Path: "/1/installations/:objectId"},

	// Cloud Functions
	{Method: "POST", Path: "/1/functions"},
}

// Static
func BenchmarkNewParse(b *testing.B) {
	benchs := []struct {
		name   string
		method string
		url    string
	}{
		{"static", "GET", "/1/users"},
		{"param/1", "GET", "/1/classes/go"},
		{"param/2", "GET", "/1/classes/go/123456789"},
	}

	for name, builder := range router.GetRegistry() {
		b.Run(name, func(b *testing.B) {
			router := builder.Build(parseAPI, router.SkipDataMode)

			for _, bench := range benchs {
				req, _ := http.NewRequest(bench.method, bench.url, nil)
				b.Run(bench.name, func(b *testing.B) {
					benchRequest(b, router, req)
				})
			}

			b.Run("all", func(b *testing.B) {
				benchRoutes(b, router, parseAPI)
			})
		})
	}
}
