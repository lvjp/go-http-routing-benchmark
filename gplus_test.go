// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

// Google+
// https://developers.google.com/+/api/latest/
// (in reality this is just a subset of a much larger API)
var gplusAPI = []router.Route{
	// People
	{Method: "GET", Path: "/people/:userId"},
	{Method: "GET", Path: "/people"},
	{Method: "GET", Path: "/activities/:activityId/people/:collection"},
	{Method: "GET", Path: "/people/:userId/people/:collection"},
	{Method: "GET", Path: "/people/:userId/openIdConnect"},

	// Activities
	{Method: "GET", Path: "/people/:userId/activities/:collection"},
	{Method: "GET", Path: "/activities/:activityId"},
	{Method: "GET", Path: "/activities"},

	// Comments
	{Method: "GET", Path: "/activities/:activityId/comments"},
	{Method: "GET", Path: "/comments/:commentId"},

	// Moments
	{Method: "POST", Path: "/people/:userId/moments/:collection"},
	{Method: "GET", Path: "/people/:userId/moments/:collection"},
	{Method: "DELETE", Path: "/moments/:id"},
}

// Static
func BenchmarkGPlus(b *testing.B) {
	benchs := []struct {
		name   string
		method string
		url    string
	}{
		{"static", "GET", "/people"},
		{"param/1", "GET", "/people/118051310819094153327"},
		{"param/2", "GET", "/people/118051310819094153327/activities/123456789"},
	}

	for name, builder := range router.GetRegistry() {
		b.Run(name, func(b *testing.B) {
			router := builder.Build(gplusAPI, router.SkipDataMode)

			for _, bench := range benchs {
				req, _ := http.NewRequest(bench.method, bench.url, nil)
				b.Run(bench.name, func(b *testing.B) {
					benchRequest(b, router, req)
				})
			}

			b.Run("all", func(b *testing.B) {
				benchRoutes(b, router, gplusAPI)
			})
		})
	}
}
