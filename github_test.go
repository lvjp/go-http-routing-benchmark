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

var (
	githubAce             http.Handler
	githubAero            http.Handler
	githubBear            http.Handler
	githubBeego           http.Handler
	githubBone            http.Handler
	githubChi             http.Handler
	githubCloudyKitRouter http.Handler
	githubDenco           http.Handler
	githubEcho            http.Handler
	githubGin             http.Handler
	githubGocraftWeb      http.Handler
	githubGoji            http.Handler
	githubGojiv2          http.Handler
	githubGoJsonRest      http.Handler
	githubGoRestful       http.Handler
	githubGorillaMux      http.Handler
	githubGowwwRouter     http.Handler
	githubHttpRouter      http.Handler
	githubHttpTreeMux     http.Handler
	githubKocha           http.Handler
	githubLARS            http.Handler
	githubMacaron         http.Handler
	githubMartini         http.Handler
	githubPat             http.Handler
	githubPossum          http.Handler
	githubR2router        http.Handler
	githubRevel           http.Handler
	githubRivet           http.Handler
	githubTango           http.Handler
	githubTigerTonic      http.Handler
	githubTraffic         http.Handler
	githubVulcan          http.Handler
)

func init() {
	println("#GithubAPI Routes:", len(githubAPI))

	calcMem("Ace", func() {
		githubAce = loadAce(githubAPI)
	})
	calcMem("Aero", func() {
		githubAero = loadAero(githubAPI)
	})
	calcMem("Bear", func() {
		githubBear = loadBear(githubAPI)
	})
	calcMem("Beego", func() {
		githubBeego = loadBeego(githubAPI)
	})
	calcMem("Bone", func() {
		githubBone = loadBone(githubAPI)
	})
	calcMem("Chi", func() {
		githubChi = loadChi(githubAPI)
	})
	calcMem("CloudyKitRouter", func() {
		githubCloudyKitRouter = loadCloudyKitRouter(githubAPI)
	})
	calcMem("Denco", func() {
		githubDenco = loadDenco(githubAPI)
	})
	calcMem("Echo", func() {
		githubEcho = loadEcho(githubAPI)
	})
	calcMem("Gin", func() {
		githubGin = loadGin(githubAPI)
	})
	calcMem("GocraftWeb", func() {
		githubGocraftWeb = loadGocraftWeb(githubAPI)
	})
	calcMem("Goji", func() {
		githubGoji = loadGoji(githubAPI)
	})
	calcMem("Gojiv2", func() {
		githubGojiv2 = loadGojiv2(githubAPI)
	})
	calcMem("GoJsonRest", func() {
		githubGoJsonRest = loadGoJsonRest(githubAPI)
	})
	calcMem("GoRestful", func() {
		githubGoRestful = loadGoRestful(githubAPI)
	})
	calcMem("GorillaMux", func() {
		githubGorillaMux = loadGorillaMux(githubAPI)
	})
	calcMem("GowwwRouter", func() {
		githubGowwwRouter = loadGowwwRouter(githubAPI)
	})
	calcMem("HttpRouter", func() {
		githubHttpRouter = loadHttpRouter(githubAPI)
	})
	calcMem("HttpTreeMux", func() {
		githubHttpTreeMux = loadHttpTreeMux(githubAPI)
	})
	calcMem("Kocha", func() {
		githubKocha = loadKocha(githubAPI)
	})
	calcMem("LARS", func() {
		githubLARS = loadLARS(githubAPI)
	})
	calcMem("Macaron", func() {
		githubMacaron = loadMacaron(githubAPI)
	})
	calcMem("Martini", func() {
		githubMartini = loadMartini(githubAPI)
	})
	calcMem("Pat", func() {
		githubPat = loadPat(githubAPI)
	})
	calcMem("Possum", func() {
		githubPossum = loadPossum(githubAPI)
	})
	calcMem("R2router", func() {
		githubR2router = loadR2router(githubAPI)
	})
	// calcMem("Revel", func() {
	// 	githubRevel = loadRevel(githubAPI)
	// })
	calcMem("Rivet", func() {
		githubRivet = loadRivet(githubAPI)
	})
	calcMem("Tango", func() {
		githubTango = loadTango(githubAPI)
	})
	calcMem("TigerTonic", func() {
		githubTigerTonic = loadTigerTonic(githubAPI)
	})
	calcMem("Traffic", func() {
		githubTraffic = loadTraffic(githubAPI)
	})
	calcMem("Vulcan", func() {
		githubVulcan = loadVulcan(githubAPI)
	})

	println()
}

// Static
func BenchmarkAce_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubAce, req)
}
func BenchmarkAero_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubAero, req)
}
func BenchmarkBear_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubBear, req)
}
func BenchmarkBeego_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubBeego, req)
}
func BenchmarkBone_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubBone, req)
}
func BenchmarkCloudyKitRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubCloudyKitRouter, req)
}
func BenchmarkChi_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubChi, req)
}
func BenchmarkDenco_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubDenco, req)
}
func BenchmarkEcho_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubEcho, req)
}
func BenchmarkGin_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGin, req)
}
func BenchmarkGocraftWeb_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGoji_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGoji, req)
}
func BenchmarkGojiv2_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGojiv2, req)
}
func BenchmarkGoRestful_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGoRestful, req)
}
func BenchmarkGoJsonRest_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGoJsonRest, req)
}
func BenchmarkGorillaMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkGowwwRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubGowwwRouter, req)
}
func BenchmarkHttpRouter_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkKocha_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubKocha, req)
}
func BenchmarkLARS_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubLARS, req)
}
func BenchmarkMacaron_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubMacaron, req)
}
func BenchmarkMartini_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkPossum_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubPossum, req)
}
func BenchmarkR2router_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubR2router, req)
}

// func BenchmarkRevel_GithubStatic(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/user/repos", nil)
// 	benchRequest(b, githubRevel, req)
// }
func BenchmarkRivet_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubRivet, req)
}
func BenchmarkTango_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTango, req)
}
func BenchmarkTigerTonic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkVulcan_GithubStatic(b *testing.B) {
	req, _ := http.NewRequest("GET", "/user/repos", nil)
	benchRequest(b, githubVulcan, req)
}

// Param
func BenchmarkAce_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubAce, req)
}
func BenchmarkAero_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubAero, req)
}
func BenchmarkBear_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBear, req)
}
func BenchmarkBeego_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBeego, req)
}
func BenchmarkBone_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubBone, req)
}
func BenchmarkChi_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubChi, req)
}
func BenchmarkCloudyKitRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubCloudyKitRouter, req)
}
func BenchmarkDenco_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubDenco, req)
}
func BenchmarkEcho_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubEcho, req)
}
func BenchmarkGin_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGin, req)
}
func BenchmarkGocraftWeb_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGocraftWeb, req)
}
func BenchmarkGoji_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoji, req)
}
func BenchmarkGojiv2_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGojiv2, req)
}
func BenchmarkGoJsonRest_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoJsonRest, req)
}
func BenchmarkGoRestful_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGoRestful, req)
}
func BenchmarkGorillaMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGorillaMux, req)
}
func BenchmarkGowwwRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubGowwwRouter, req)
}
func BenchmarkHttpRouter_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpRouter, req)
}
func BenchmarkHttpTreeMux_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubHttpTreeMux, req)
}
func BenchmarkKocha_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubKocha, req)
}
func BenchmarkLARS_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubLARS, req)
}
func BenchmarkMacaron_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMacaron, req)
}
func BenchmarkMartini_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubMartini, req)
}
func BenchmarkPat_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubPat, req)
}
func BenchmarkPossum_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubPossum, req)
}
func BenchmarkR2router_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubR2router, req)
}

// func BenchmarkRevel_GithubParam(b *testing.B) {
// 	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
// 	benchRequest(b, githubRevel, req)
// }
func BenchmarkRivet_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubRivet, req)
}
func BenchmarkTango_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTango, req)
}
func BenchmarkTigerTonic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTigerTonic, req)
}
func BenchmarkTraffic_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubTraffic, req)
}
func BenchmarkVulcan_GithubParam(b *testing.B) {
	req, _ := http.NewRequest("GET", "/repos/julienschmidt/httprouter/stargazers", nil)
	benchRequest(b, githubVulcan, req)
}

// All routes
func BenchmarkAce_GithubAll(b *testing.B) {
	benchRoutes(b, githubAce, githubAPI)
}
func BenchmarkAero_GithubAll(b *testing.B) {
	benchRoutes(b, githubAero, githubAPI)
}
func BenchmarkBear_GithubAll(b *testing.B) {
	benchRoutes(b, githubBear, githubAPI)
}
func BenchmarkBeego_GithubAll(b *testing.B) {
	benchRoutes(b, githubBeego, githubAPI)
}
func BenchmarkBone_GithubAll(b *testing.B) {
	benchRoutes(b, githubBone, githubAPI)
}
func BenchmarkChi_GithubAll(b *testing.B) {
	benchRoutes(b, githubChi, githubAPI)
}
func BenchmarkCloudyKitRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubCloudyKitRouter, githubAPI)
}
func BenchmarkDenco_GithubAll(b *testing.B) {
	benchRoutes(b, githubDenco, githubAPI)
}
func BenchmarkEcho_GithubAll(b *testing.B) {
	benchRoutes(b, githubEcho, githubAPI)
}
func BenchmarkGin_GithubAll(b *testing.B) {
	benchRoutes(b, githubGin, githubAPI)
}
func BenchmarkGocraftWeb_GithubAll(b *testing.B) {
	benchRoutes(b, githubGocraftWeb, githubAPI)
}
func BenchmarkGoji_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoji, githubAPI)
}
func BenchmarkGojiv2_GithubAll(b *testing.B) {
	benchRoutes(b, githubGojiv2, githubAPI)
}
func BenchmarkGoJsonRest_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoJsonRest, githubAPI)
}
func BenchmarkGoRestful_GithubAll(b *testing.B) {
	benchRoutes(b, githubGoRestful, githubAPI)
}
func BenchmarkGorillaMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubGorillaMux, githubAPI)
}
func BenchmarkGowwwRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubGowwwRouter, githubAPI)
}
func BenchmarkHttpRouter_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpRouter, githubAPI)
}
func BenchmarkHttpTreeMux_GithubAll(b *testing.B) {
	benchRoutes(b, githubHttpTreeMux, githubAPI)
}
func BenchmarkKocha_GithubAll(b *testing.B) {
	benchRoutes(b, githubKocha, githubAPI)
}
func BenchmarkLARS_GithubAll(b *testing.B) {
	benchRoutes(b, githubLARS, githubAPI)
}
func BenchmarkMacaron_GithubAll(b *testing.B) {
	benchRoutes(b, githubMacaron, githubAPI)
}
func BenchmarkMartini_GithubAll(b *testing.B) {
	benchRoutes(b, githubMartini, githubAPI)
}
func BenchmarkPat_GithubAll(b *testing.B) {
	benchRoutes(b, githubPat, githubAPI)
}
func BenchmarkPossum_GithubAll(b *testing.B) {
	benchRoutes(b, githubPossum, githubAPI)
}
func BenchmarkR2router_GithubAll(b *testing.B) {
	benchRoutes(b, githubR2router, githubAPI)
}

// func BenchmarkRevel_GithubAll(b *testing.B) {
// 	benchRoutes(b, githubRevel, githubAPI)
// }
func BenchmarkRivet_GithubAll(b *testing.B) {
	benchRoutes(b, githubRivet, githubAPI)
}
func BenchmarkTango_GithubAll(b *testing.B) {
	benchRoutes(b, githubTango, githubAPI)
}
func BenchmarkTigerTonic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTigerTonic, githubAPI)
}
func BenchmarkTraffic_GithubAll(b *testing.B) {
	benchRoutes(b, githubTraffic, githubAPI)
}
func BenchmarkVulcan_GithubAll(b *testing.B) {
	benchRoutes(b, githubVulcan, githubAPI)
}
