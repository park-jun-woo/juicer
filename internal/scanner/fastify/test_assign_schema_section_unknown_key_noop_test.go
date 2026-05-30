//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestAssignSchemaSection_UnknownKeyNoop 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestAssignSchemaSection_UnknownKeyNoop(t *testing.T) {

	pairs, src := schemaPairs(t, `{ headers: { type: "object" } }`)
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	for _, p := range pairs {
		assignSchemaSection(si, p, src)
	}
	if si.Body != nil || si.Querystring != nil || si.Params != nil || len(si.Response) != 0 {
		t.Fatalf("unknown key should be a no-op, got %v", si)
	}
}
