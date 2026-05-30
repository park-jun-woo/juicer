//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestBuildRequest_SchemaPresentButNoSchemaKey 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildRequest_SchemaPresentButNoSchemaKey(t *testing.T) {

	obj, src := firstObject(t, `{ config: { rateLimit: true } }`)
	r := routeInfo{Method: "GET", Schema: obj}
	req, has := buildRequest(r, nil, src, map[string]*sitter.Node{})
	if has || req.Body != nil {
		t.Fatalf("expected empty request when schema info nil, got has=%v", has)
	}
}
