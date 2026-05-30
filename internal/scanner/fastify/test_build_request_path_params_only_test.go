//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildRequest_PathParamsOnly 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildRequest_PathParamsOnly(t *testing.T) {
	r := routeInfo{Method: "GET"}
	req, has := buildRequest(r, []string{"id"}, []byte(""), map[string]*sitter.Node{})
	if !has || len(req.PathParams) != 1 {
		t.Fatalf("expected path params, got has=%v %v", has, req.PathParams)
	}
}
