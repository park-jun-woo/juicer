//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildRequest_NoSchemaNoParams 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildRequest_NoSchemaNoParams(t *testing.T) {
	r := routeInfo{Method: "GET", Schema: nil}
	req, has := buildRequest(r, nil, []byte(""), map[string]*sitter.Node{})
	if has || req.Body != nil {
		t.Fatalf("expected empty request, got has=%v", has)
	}
}
