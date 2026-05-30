//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractOneRoute 테스트
package quarkus

import "testing"

func TestExtractOneRoute(t *testing.T) {
	fi := qFileInfo(t, `class R {
		@GET @Path("/{id}") public UserDto get(@PathParam("id") Long id) { return null; }
	}`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep, ok := extractOneRoute(m, fi)
	if !ok {
		t.Fatal("expected route")
	}
	if ep.method != "GET" || ep.path != "/{id}" || ep.handler != "get" {
		t.Fatalf("got %+v", ep)
	}
	if len(ep.params) != 1 {
		t.Fatalf("params: %+v", ep.params)
	}
}
