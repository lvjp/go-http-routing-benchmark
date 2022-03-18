// Copyright 2013 Julien Schmidt. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package main

import (
	"testing"

	"github.com/lvjp/go-http-routing-benchmark/router"
)

var staticRoutes = []router.Route{
	{Method: "GET", Path: "/"},
	{Method: "GET", Path: "/cmd.html"},
	{Method: "GET", Path: "/code.html"},
	{Method: "GET", Path: "/contrib.html"},
	{Method: "GET", Path: "/contribute.html"},
	{Method: "GET", Path: "/debugging_with_gdb.html"},
	{Method: "GET", Path: "/docs.html"},
	{Method: "GET", Path: "/effective_go.html"},
	{Method: "GET", Path: "/files.log"},
	{Method: "GET", Path: "/gccgo_contribute.html"},
	{Method: "GET", Path: "/gccgo_install.html"},
	{Method: "GET", Path: "/go-logo-black.png"},
	{Method: "GET", Path: "/go-logo-blue.png"},
	{Method: "GET", Path: "/go-logo-white.png"},
	{Method: "GET", Path: "/go1.1.html"},
	{Method: "GET", Path: "/go1.2.html"},
	{Method: "GET", Path: "/go1.html"},
	{Method: "GET", Path: "/go1compat.html"},
	{Method: "GET", Path: "/go_faq.html"},
	{Method: "GET", Path: "/go_mem.html"},
	{Method: "GET", Path: "/go_spec.html"},
	{Method: "GET", Path: "/help.html"},
	{Method: "GET", Path: "/ie.css"},
	{Method: "GET", Path: "/install-source.html"},
	{Method: "GET", Path: "/install.html"},
	{Method: "GET", Path: "/logo-153x55.png"},
	{Method: "GET", Path: "/Makefile"},
	{Method: "GET", Path: "/root.html"},
	{Method: "GET", Path: "/share.png"},
	{Method: "GET", Path: "/sieve.gif"},
	{Method: "GET", Path: "/tos.html"},
	{Method: "GET", Path: "/articles"},
	{Method: "GET", Path: "/articles/go_command.html"},
	{Method: "GET", Path: "/articles/index.html"},
	{Method: "GET", Path: "/articles/wiki"},
	{Method: "GET", Path: "/articles/wiki/edit.html"},
	{Method: "GET", Path: "/articles/wiki/final-noclosure.go"},
	{Method: "GET", Path: "/articles/wiki/final-noerror.go"},
	{Method: "GET", Path: "/articles/wiki/final-parsetemplate.go"},
	{Method: "GET", Path: "/articles/wiki/final-template.go"},
	{Method: "GET", Path: "/articles/wiki/final.go"},
	{Method: "GET", Path: "/articles/wiki/get.go"},
	{Method: "GET", Path: "/articles/wiki/http-sample.go"},
	{Method: "GET", Path: "/articles/wiki/index.html"},
	{Method: "GET", Path: "/articles/wiki/Makefile"},
	{Method: "GET", Path: "/articles/wiki/notemplate.go"},
	{Method: "GET", Path: "/articles/wiki/part1-noerror.go"},
	{Method: "GET", Path: "/articles/wiki/part1.go"},
	{Method: "GET", Path: "/articles/wiki/part2.go"},
	{Method: "GET", Path: "/articles/wiki/part3-errorhandling.go"},
	{Method: "GET", Path: "/articles/wiki/part3.go"},
	{Method: "GET", Path: "/articles/wiki/test.bash"},
	{Method: "GET", Path: "/articles/wiki/test_edit.good"},
	{Method: "GET", Path: "/articles/wiki/test_Test.txt.good"},
	{Method: "GET", Path: "/articles/wiki/test_view.good"},
	{Method: "GET", Path: "/articles/wiki/view.html"},
	{Method: "GET", Path: "/codewalk"},
	{Method: "GET", Path: "/codewalk/codewalk.css"},
	{Method: "GET", Path: "/codewalk/codewalk.js"},
	{Method: "GET", Path: "/codewalk/codewalk.xml"},
	{Method: "GET", Path: "/codewalk/functions.xml"},
	{Method: "GET", Path: "/codewalk/markov.go"},
	{Method: "GET", Path: "/codewalk/markov.xml"},
	{Method: "GET", Path: "/codewalk/pig.go"},
	{Method: "GET", Path: "/codewalk/popout.png"},
	{Method: "GET", Path: "/codewalk/run"},
	{Method: "GET", Path: "/codewalk/sharemem.xml"},
	{Method: "GET", Path: "/codewalk/urlpoll.go"},
	{Method: "GET", Path: "/devel"},
	{Method: "GET", Path: "/devel/release.html"},
	{Method: "GET", Path: "/devel/weekly.html"},
	{Method: "GET", Path: "/gopher"},
	{Method: "GET", Path: "/gopher/appenginegopher.jpg"},
	{Method: "GET", Path: "/gopher/appenginegophercolor.jpg"},
	{Method: "GET", Path: "/gopher/appenginelogo.gif"},
	{Method: "GET", Path: "/gopher/bumper.png"},
	{Method: "GET", Path: "/gopher/bumper192x108.png"},
	{Method: "GET", Path: "/gopher/bumper320x180.png"},
	{Method: "GET", Path: "/gopher/bumper480x270.png"},
	{Method: "GET", Path: "/gopher/bumper640x360.png"},
	{Method: "GET", Path: "/gopher/doc.png"},
	{Method: "GET", Path: "/gopher/frontpage.png"},
	{Method: "GET", Path: "/gopher/gopherbw.png"},
	{Method: "GET", Path: "/gopher/gophercolor.png"},
	{Method: "GET", Path: "/gopher/gophercolor16x16.png"},
	{Method: "GET", Path: "/gopher/help.png"},
	{Method: "GET", Path: "/gopher/pkg.png"},
	{Method: "GET", Path: "/gopher/project.png"},
	{Method: "GET", Path: "/gopher/ref.png"},
	{Method: "GET", Path: "/gopher/run.png"},
	{Method: "GET", Path: "/gopher/talks.png"},
	{Method: "GET", Path: "/gopher/pencil"},
	{Method: "GET", Path: "/gopher/pencil/gopherhat.jpg"},
	{Method: "GET", Path: "/gopher/pencil/gopherhelmet.jpg"},
	{Method: "GET", Path: "/gopher/pencil/gophermega.jpg"},
	{Method: "GET", Path: "/gopher/pencil/gopherrunning.jpg"},
	{Method: "GET", Path: "/gopher/pencil/gopherswim.jpg"},
	{Method: "GET", Path: "/gopher/pencil/gopherswrench.jpg"},
	{Method: "GET", Path: "/play"},
	{Method: "GET", Path: "/play/fib.go"},
	{Method: "GET", Path: "/play/hello.go"},
	{Method: "GET", Path: "/play/life.go"},
	{Method: "GET", Path: "/play/peano.go"},
	{Method: "GET", Path: "/play/pi.go"},
	{Method: "GET", Path: "/play/sieve.go"},
	{Method: "GET", Path: "/play/solitaire.go"},
	{Method: "GET", Path: "/play/tree.go"},
	{Method: "GET", Path: "/progs"},
	{Method: "GET", Path: "/progs/cgo1.go"},
	{Method: "GET", Path: "/progs/cgo2.go"},
	{Method: "GET", Path: "/progs/cgo3.go"},
	{Method: "GET", Path: "/progs/cgo4.go"},
	{Method: "GET", Path: "/progs/defer.go"},
	{Method: "GET", Path: "/progs/defer.out"},
	{Method: "GET", Path: "/progs/defer2.go"},
	{Method: "GET", Path: "/progs/defer2.out"},
	{Method: "GET", Path: "/progs/eff_bytesize.go"},
	{Method: "GET", Path: "/progs/eff_bytesize.out"},
	{Method: "GET", Path: "/progs/eff_qr.go"},
	{Method: "GET", Path: "/progs/eff_sequence.go"},
	{Method: "GET", Path: "/progs/eff_sequence.out"},
	{Method: "GET", Path: "/progs/eff_unused1.go"},
	{Method: "GET", Path: "/progs/eff_unused2.go"},
	{Method: "GET", Path: "/progs/error.go"},
	{Method: "GET", Path: "/progs/error2.go"},
	{Method: "GET", Path: "/progs/error3.go"},
	{Method: "GET", Path: "/progs/error4.go"},
	{Method: "GET", Path: "/progs/go1.go"},
	{Method: "GET", Path: "/progs/gobs1.go"},
	{Method: "GET", Path: "/progs/gobs2.go"},
	{Method: "GET", Path: "/progs/image_draw.go"},
	{Method: "GET", Path: "/progs/image_package1.go"},
	{Method: "GET", Path: "/progs/image_package1.out"},
	{Method: "GET", Path: "/progs/image_package2.go"},
	{Method: "GET", Path: "/progs/image_package2.out"},
	{Method: "GET", Path: "/progs/image_package3.go"},
	{Method: "GET", Path: "/progs/image_package3.out"},
	{Method: "GET", Path: "/progs/image_package4.go"},
	{Method: "GET", Path: "/progs/image_package4.out"},
	{Method: "GET", Path: "/progs/image_package5.go"},
	{Method: "GET", Path: "/progs/image_package5.out"},
	{Method: "GET", Path: "/progs/image_package6.go"},
	{Method: "GET", Path: "/progs/image_package6.out"},
	{Method: "GET", Path: "/progs/interface.go"},
	{Method: "GET", Path: "/progs/interface2.go"},
	{Method: "GET", Path: "/progs/interface2.out"},
	{Method: "GET", Path: "/progs/json1.go"},
	{Method: "GET", Path: "/progs/json2.go"},
	{Method: "GET", Path: "/progs/json2.out"},
	{Method: "GET", Path: "/progs/json3.go"},
	{Method: "GET", Path: "/progs/json4.go"},
	{Method: "GET", Path: "/progs/json5.go"},
	{Method: "GET", Path: "/progs/run"},
	{Method: "GET", Path: "/progs/slices.go"},
	{Method: "GET", Path: "/progs/timeout1.go"},
	{Method: "GET", Path: "/progs/timeout2.go"},
	{Method: "GET", Path: "/progs/update.bash"},
}

// All routes
func BenchmarkStaticAll(b *testing.B) {
	for name, builder := range router.GetRegistry() {
		router := builder.Build(staticRoutes, router.SkipDataMode)
		b.Run(name, func(b *testing.B) {
			benchRoutes(b, router, staticRoutes)
		})
	}
}
