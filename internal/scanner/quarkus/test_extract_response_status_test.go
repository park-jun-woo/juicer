//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractResponseStatus 테스트
package quarkus

import "testing"

func TestExtractResponseStatus(t *testing.T) {
	fi := qFileInfo(t, `class R {
		@POST public Response create() { return Response.status(201).build(); }
	}`)
	m := findAllByType(fi.root, "method_declaration")[0]
	ep := &endpointInfo{}
	extractResponseStatus(m, fi.src, ep)
	if ep.statusCode != "201" {
		t.Fatalf("got %q", ep.statusCode)
	}
}
