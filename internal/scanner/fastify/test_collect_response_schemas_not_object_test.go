//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestCollectResponseSchemas_NotObject 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectResponseSchemas_NotObject(t *testing.T) {

	fi := mustParse(t, []byte("const x = [1, 2];\n"))
	arr := findAllByType(fi.Root, "array")[0]
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	collectResponseSchemas(si, arr, fi.Src)
	if len(si.Response) != 0 {
		t.Fatalf("expected no responses for non-object, got %d", len(si.Response))
	}
}
