// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"net/http"
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

// http://developer.github.com/v3/
var githubAPI = []router.Route{
	// OAuth Authorizations
	{Method: "GET", Path: "/authorizations"},
	{Method: "GET", Path: "/authorizations/:id"},
	{Method: "POST", Path: "/authorizations"},
	//{"PUT", "/authorizations/clients/:client_id"},
	//{"PATCH", "/authorizations/:id"},
	{Method: "DELETE", Path: "/authorizations/:id"},
	{Method: "GET", Path: "/applications/:client_id/tokens/:access_token"},
	{Method: "DELETE", Path: "/applications/:client_id/tokens"},
	{Method: "DELETE", Path: "/applications/:client_id/tokens/:access_token"},

	// Activity
	{Method: "GET", Path: "/events"},
	{Method: "GET", Path: "/repos/:owner/:repo/events"},
	{Method: "GET", Path: "/networks/:owner/:repo/events"},
	{Method: "GET", Path: "/orgs/:org/events"},
	{Method: "GET", Path: "/users/:user/received_events"},
	{Method: "GET", Path: "/users/:user/received_events/public"},
	{Method: "GET", Path: "/users/:user/events"},
	{Method: "GET", Path: "/users/:user/events/public"},
	{Method: "GET", Path: "/users/:user/events/orgs/:org"},
	{Method: "GET", Path: "/feeds"},
	{Method: "GET", Path: "/notifications"},
	{Method: "GET", Path: "/repos/:owner/:repo/notifications"},
	{Method: "PUT", Path: "/notifications"},
	{Method: "PUT", Path: "/repos/:owner/:repo/notifications"},
	{Method: "GET", Path: "/notifications/threads/:id"},
	//{"PATCH", "/notifications/threads/:id"},
	{Method: "GET", Path: "/notifications/threads/:id/subscription"},
	{Method: "PUT", Path: "/notifications/threads/:id/subscription"},
	{Method: "DELETE", Path: "/notifications/threads/:id/subscription"},
	{Method: "GET", Path: "/repos/:owner/:repo/stargazers"},
	{Method: "GET", Path: "/users/:user/starred"},
	{Method: "GET", Path: "/user/starred"},
	{Method: "GET", Path: "/user/starred/:owner/:repo"},
	{Method: "PUT", Path: "/user/starred/:owner/:repo"},
	{Method: "DELETE", Path: "/user/starred/:owner/:repo"},
	{Method: "GET", Path: "/repos/:owner/:repo/subscribers"},
	{Method: "GET", Path: "/users/:user/subscriptions"},
	{Method: "GET", Path: "/user/subscriptions"},
	{Method: "GET", Path: "/repos/:owner/:repo/subscription"},
	{Method: "PUT", Path: "/repos/:owner/:repo/subscription"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/subscription"},
	{Method: "GET", Path: "/user/subscriptions/:owner/:repo"},
	{Method: "PUT", Path: "/user/subscriptions/:owner/:repo"},
	{Method: "DELETE", Path: "/user/subscriptions/:owner/:repo"},

	// Gists
	{Method: "GET", Path: "/users/:user/gists"},
	{Method: "GET", Path: "/gists"},
	//{"GET", "/gists/public"},
	//{"GET", "/gists/starred"},
	{Method: "GET", Path: "/gists/:id"},
	{Method: "POST", Path: "/gists"},
	//{"PATCH", "/gists/:id"},
	{Method: "PUT", Path: "/gists/:id/star"},
	{Method: "DELETE", Path: "/gists/:id/star"},
	{Method: "GET", Path: "/gists/:id/star"},
	{Method: "POST", Path: "/gists/:id/forks"},
	{Method: "DELETE", Path: "/gists/:id"},

	// Git Data
	{Method: "GET", Path: "/repos/:owner/:repo/git/blobs/:sha"},
	{Method: "POST", Path: "/repos/:owner/:repo/git/blobs"},
	{Method: "GET", Path: "/repos/:owner/:repo/git/commits/:sha"},
	{Method: "POST", Path: "/repos/:owner/:repo/git/commits"},
	//{"GET", "/repos/:owner/:repo/git/refs/*ref"},
	{Method: "GET", Path: "/repos/:owner/:repo/git/refs"},
	{Method: "POST", Path: "/repos/:owner/:repo/git/refs"},
	//{"PATCH", "/repos/:owner/:repo/git/refs/*ref"},
	//{"DELETE", "/repos/:owner/:repo/git/refs/*ref"},
	{Method: "GET", Path: "/repos/:owner/:repo/git/tags/:sha"},
	{Method: "POST", Path: "/repos/:owner/:repo/git/tags"},
	{Method: "GET", Path: "/repos/:owner/:repo/git/trees/:sha"},
	{Method: "POST", Path: "/repos/:owner/:repo/git/trees"},

	// Issues
	{Method: "GET", Path: "/issues"},
	{Method: "GET", Path: "/user/issues"},
	{Method: "GET", Path: "/orgs/:org/issues"},
	{Method: "GET", Path: "/repos/:owner/:repo/issues"},
	{Method: "GET", Path: "/repos/:owner/:repo/issues/:number"},
	{Method: "POST", Path: "/repos/:owner/:repo/issues"},
	//{"PATCH", "/repos/:owner/:repo/issues/:number"},
	{Method: "GET", Path: "/repos/:owner/:repo/assignees"},
	{Method: "GET", Path: "/repos/:owner/:repo/assignees/:assignee"},
	{Method: "GET", Path: "/repos/:owner/:repo/issues/:number/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments"},
	//{"GET", "/repos/:owner/:repo/issues/comments/:id"},
	{Method: "POST", Path: "/repos/:owner/:repo/issues/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/issues/comments/:id"},
	//{"DELETE", "/repos/:owner/:repo/issues/comments/:id"},
	{Method: "GET", Path: "/repos/:owner/:repo/issues/:number/events"},
	//{"GET", "/repos/:owner/:repo/issues/events"},
	//{"GET", "/repos/:owner/:repo/issues/events/:id"},
	{Method: "GET", Path: "/repos/:owner/:repo/labels"},
	{Method: "GET", Path: "/repos/:owner/:repo/labels/:name"},
	{Method: "POST", Path: "/repos/:owner/:repo/labels"},
	//{"PATCH", "/repos/:owner/:repo/labels/:name"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/labels/:name"},
	{Method: "GET", Path: "/repos/:owner/:repo/issues/:number/labels"},
	{Method: "POST", Path: "/repos/:owner/:repo/issues/:number/labels"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/issues/:number/labels/:name"},
	{Method: "PUT", Path: "/repos/:owner/:repo/issues/:number/labels"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/issues/:number/labels"},
	{Method: "GET", Path: "/repos/:owner/:repo/milestones/:number/labels"},
	{Method: "GET", Path: "/repos/:owner/:repo/milestones"},
	{Method: "GET", Path: "/repos/:owner/:repo/milestones/:number"},
	{Method: "POST", Path: "/repos/:owner/:repo/milestones"},
	//{"PATCH", "/repos/:owner/:repo/milestones/:number"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/milestones/:number"},

	// Miscellaneous
	{Method: "GET", Path: "/emojis"},
	{Method: "GET", Path: "/gitignore/templates"},
	{Method: "GET", Path: "/gitignore/templates/:name"},
	{Method: "POST", Path: "/markdown"},
	{Method: "POST", Path: "/markdown/raw"},
	{Method: "GET", Path: "/meta"},
	{Method: "GET", Path: "/rate_limit"},

	// Organizations
	{Method: "GET", Path: "/users/:user/orgs"},
	{Method: "GET", Path: "/user/orgs"},
	{Method: "GET", Path: "/orgs/:org"},
	//{"PATCH", "/orgs/:org"},
	{Method: "GET", Path: "/orgs/:org/members"},
	{Method: "GET", Path: "/orgs/:org/members/:user"},
	{Method: "DELETE", Path: "/orgs/:org/members/:user"},
	{Method: "GET", Path: "/orgs/:org/public_members"},
	{Method: "GET", Path: "/orgs/:org/public_members/:user"},
	{Method: "PUT", Path: "/orgs/:org/public_members/:user"},
	{Method: "DELETE", Path: "/orgs/:org/public_members/:user"},
	{Method: "GET", Path: "/orgs/:org/teams"},
	{Method: "GET", Path: "/teams/:id"},
	{Method: "POST", Path: "/orgs/:org/teams"},
	//{"PATCH", "/teams/:id"},
	{Method: "DELETE", Path: "/teams/:id"},
	{Method: "GET", Path: "/teams/:id/members"},
	{Method: "GET", Path: "/teams/:id/members/:user"},
	{Method: "PUT", Path: "/teams/:id/members/:user"},
	{Method: "DELETE", Path: "/teams/:id/members/:user"},
	{Method: "GET", Path: "/teams/:id/repos"},
	{Method: "GET", Path: "/teams/:id/repos/:owner/:repo"},
	{Method: "PUT", Path: "/teams/:id/repos/:owner/:repo"},
	{Method: "DELETE", Path: "/teams/:id/repos/:owner/:repo"},
	{Method: "GET", Path: "/user/teams"},

	// Pull Requests
	{Method: "GET", Path: "/repos/:owner/:repo/pulls"},
	{Method: "GET", Path: "/repos/:owner/:repo/pulls/:number"},
	{Method: "POST", Path: "/repos/:owner/:repo/pulls"},
	//{"PATCH", "/repos/:owner/:repo/pulls/:number"},
	{Method: "GET", Path: "/repos/:owner/:repo/pulls/:number/commits"},
	{Method: "GET", Path: "/repos/:owner/:repo/pulls/:number/files"},
	{Method: "GET", Path: "/repos/:owner/:repo/pulls/:number/merge"},
	{Method: "PUT", Path: "/repos/:owner/:repo/pulls/:number/merge"},
	{Method: "GET", Path: "/repos/:owner/:repo/pulls/:number/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments"},
	//{"GET", "/repos/:owner/:repo/pulls/comments/:number"},
	{Method: "PUT", Path: "/repos/:owner/:repo/pulls/:number/comments"},
	//{"PATCH", "/repos/:owner/:repo/pulls/comments/:number"},
	//{"DELETE", "/repos/:owner/:repo/pulls/comments/:number"},

	// Repositories
	{Method: "GET", Path: "/user/repos"},
	{Method: "GET", Path: "/users/:user/repos"},
	{Method: "GET", Path: "/orgs/:org/repos"},
	{Method: "GET", Path: "/repositories"},
	{Method: "POST", Path: "/user/repos"},
	{Method: "POST", Path: "/orgs/:org/repos"},
	{Method: "GET", Path: "/repos/:owner/:repo"},
	//{"PATCH", "/repos/:owner/:repo"},
	{Method: "GET", Path: "/repos/:owner/:repo/contributors"},
	{Method: "GET", Path: "/repos/:owner/:repo/languages"},
	{Method: "GET", Path: "/repos/:owner/:repo/teams"},
	{Method: "GET", Path: "/repos/:owner/:repo/tags"},
	{Method: "GET", Path: "/repos/:owner/:repo/branches"},
	{Method: "GET", Path: "/repos/:owner/:repo/branches/:branch"},
	{Method: "DELETE", Path: "/repos/:owner/:repo"},
	{Method: "GET", Path: "/repos/:owner/:repo/collaborators"},
	{Method: "GET", Path: "/repos/:owner/:repo/collaborators/:user"},
	{Method: "PUT", Path: "/repos/:owner/:repo/collaborators/:user"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/collaborators/:user"},
	{Method: "GET", Path: "/repos/:owner/:repo/comments"},
	{Method: "GET", Path: "/repos/:owner/:repo/commits/:sha/comments"},
	{Method: "POST", Path: "/repos/:owner/:repo/commits/:sha/comments"},
	{Method: "GET", Path: "/repos/:owner/:repo/comments/:id"},
	//{"PATCH", "/repos/:owner/:repo/comments/:id"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/comments/:id"},
	{Method: "GET", Path: "/repos/:owner/:repo/commits"},
	{Method: "GET", Path: "/repos/:owner/:repo/commits/:sha"},
	{Method: "GET", Path: "/repos/:owner/:repo/readme"},
	//{"GET", "/repos/:owner/:repo/contents/*path"},
	//{"PUT", "/repos/:owner/:repo/contents/*path"},
	//{"DELETE", "/repos/:owner/:repo/contents/*path"},
	//{"GET", "/repos/:owner/:repo/:archive_format/:ref"},
	{Method: "GET", Path: "/repos/:owner/:repo/keys"},
	{Method: "GET", Path: "/repos/:owner/:repo/keys/:id"},
	{Method: "POST", Path: "/repos/:owner/:repo/keys"},
	//{"PATCH", "/repos/:owner/:repo/keys/:id"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/keys/:id"},
	{Method: "GET", Path: "/repos/:owner/:repo/downloads"},
	{Method: "GET", Path: "/repos/:owner/:repo/downloads/:id"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/downloads/:id"},
	{Method: "GET", Path: "/repos/:owner/:repo/forks"},
	{Method: "POST", Path: "/repos/:owner/:repo/forks"},
	{Method: "GET", Path: "/repos/:owner/:repo/hooks"},
	{Method: "GET", Path: "/repos/:owner/:repo/hooks/:id"},
	{Method: "POST", Path: "/repos/:owner/:repo/hooks"},
	//{"PATCH", "/repos/:owner/:repo/hooks/:id"},
	{Method: "POST", Path: "/repos/:owner/:repo/hooks/:id/tests"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/hooks/:id"},
	{Method: "POST", Path: "/repos/:owner/:repo/merges"},
	{Method: "GET", Path: "/repos/:owner/:repo/releases"},
	{Method: "GET", Path: "/repos/:owner/:repo/releases/:id"},
	{Method: "POST", Path: "/repos/:owner/:repo/releases"},
	//{"PATCH", "/repos/:owner/:repo/releases/:id"},
	{Method: "DELETE", Path: "/repos/:owner/:repo/releases/:id"},
	{Method: "GET", Path: "/repos/:owner/:repo/releases/:id/assets"},
	{Method: "GET", Path: "/repos/:owner/:repo/stats/contributors"},
	{Method: "GET", Path: "/repos/:owner/:repo/stats/commit_activity"},
	{Method: "GET", Path: "/repos/:owner/:repo/stats/code_frequency"},
	{Method: "GET", Path: "/repos/:owner/:repo/stats/participation"},
	{Method: "GET", Path: "/repos/:owner/:repo/stats/punch_card"},
	{Method: "GET", Path: "/repos/:owner/:repo/statuses/:ref"},
	{Method: "POST", Path: "/repos/:owner/:repo/statuses/:ref"},

	// Search
	{Method: "GET", Path: "/search/repositories"},
	{Method: "GET", Path: "/search/code"},
	{Method: "GET", Path: "/search/issues"},
	{Method: "GET", Path: "/search/users"},
	{Method: "GET", Path: "/legacy/issues/search/:owner/:repository/:state/:keyword"},
	{Method: "GET", Path: "/legacy/repos/search/:keyword"},
	{Method: "GET", Path: "/legacy/user/search/:keyword"},
	{Method: "GET", Path: "/legacy/user/email/:email"},

	// Users
	{Method: "GET", Path: "/users/:user"},
	{Method: "GET", Path: "/user"},
	//{"PATCH", "/user"},
	{Method: "GET", Path: "/users"},
	{Method: "GET", Path: "/user/emails"},
	{Method: "POST", Path: "/user/emails"},
	{Method: "DELETE", Path: "/user/emails"},
	{Method: "GET", Path: "/users/:user/followers"},
	{Method: "GET", Path: "/user/followers"},
	{Method: "GET", Path: "/users/:user/following"},
	{Method: "GET", Path: "/user/following"},
	{Method: "GET", Path: "/user/following/:user"},
	{Method: "GET", Path: "/users/:user/following/:target_user"},
	{Method: "PUT", Path: "/user/following/:user"},
	{Method: "DELETE", Path: "/user/following/:user"},
	{Method: "GET", Path: "/users/:user/keys"},
	{Method: "GET", Path: "/user/keys"},
	{Method: "GET", Path: "/user/keys/:id"},
	{Method: "POST", Path: "/user/keys"},
	//{"PATCH", "/user/keys/:id"},
	{Method: "DELETE", Path: "/user/keys/:id"},
}

// Static
func BenchmarkGithub(b *testing.B) {
	benchs := []struct {
		name   string
		method string
		url    string
	}{
		{"static", "GET", "/user/repos"},
		{"param", "GET", "/repos/julienschmidt/httprouter/stargazers"},
	}

	for name, builder := range router.GetRegistry() {
		router := builder.Build(githubAPI, router.SkipDataMode)

		for _, bench := range benchs {
			req, _ := http.NewRequest(bench.method, bench.url, nil)
			b.Run(name+"/"+bench.name, func(b *testing.B) {
				benchRequest(b, router, req)
			})
		}

		b.Run(name+"/all", func(b *testing.B) {
			benchRoutes(b, router, githubAPI)
		})

	}
}
