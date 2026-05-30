//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestAssignSchemaSection_AllSections 테스트
package fastify

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

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
