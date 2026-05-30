//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what assignSchemaSection 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

// schemaPairs returns the pair nodes of the first object literal.
func schemaPairs(t *testing.T, objSrc string) ([]*sitter.Node, []byte) {
	t.Helper()
	fi := mustParse(t, []byte("const x = "+objSrc+";\n"))
	objs := findAllByType(fi.Root, "object")
	if len(objs) == 0 {
		t.Fatal("no object")
	}
	return childrenOfType(objs[0], "pair"), fi.Src
}

func TestAssignSchemaSection_AllSections(t *testing.T) {
	pairs, src := schemaPairs(t, `{
  body: { type: "object" },
  querystring: { type: "object" },
  params: { type: "object" },
  response: { "200": { type: "object" } },
  unknown: 1
}`)
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	for _, p := range pairs {
		assignSchemaSection(si, p, src)
	}
	if si.Body == nil {
		t.Error("body not set")
	}
	if si.Querystring == nil {
		t.Error("querystring not set")
	}
	if si.Params == nil {
		t.Error("params not set")
	}
	if len(si.Response) == 0 {
		t.Error("response not collected")
	}
}

func TestAssignSchemaSection_UnknownKeyNoop(t *testing.T) {
	// a recognized-shape pair with an unknown key leaves si untouched (default case).
	pairs, src := schemaPairs(t, `{ headers: { type: "object" } }`)
	si := &schemaInfo{Response: make(map[string]*sitter.Node)}
	for _, p := range pairs {
		assignSchemaSection(si, p, src)
	}
	if si.Body != nil || si.Querystring != nil || si.Params != nil || len(si.Response) != 0 {
		t.Fatalf("unknown key should be a no-op, got %v", si)
	}
}

