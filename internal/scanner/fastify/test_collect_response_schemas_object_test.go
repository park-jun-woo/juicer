//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectResponseSchemas_Object 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectResponseSchemas_Object(t *testing.T) {
	obj, src := firstObject(t, `{ "200": { type: "object" }, "404": { type: "object" } }`)
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	collectResponseSchemas(si, obj, src)
	if len(si.Response) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(si.Response))
	}
	if si.Response["200"] == nil || si.Response["404"] == nil {
		t.Fatalf("missing status codes: %v", si.Response)
	}
}
