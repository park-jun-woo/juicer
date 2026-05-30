//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what collectResponseSchemas 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
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

func TestCollectResponseSchemas_NotObject(t *testing.T) {
	// val is an array, not an object -> early return, nothing collected
	fi := mustParse(t, []byte("const x = [1, 2];\n"))
	arr := findAllByType(fi.Root, "array")[0]
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	collectResponseSchemas(si, arr, fi.Src)
	if len(si.Response) != 0 {
		t.Fatalf("expected no responses for non-object, got %d", len(si.Response))
	}
}
