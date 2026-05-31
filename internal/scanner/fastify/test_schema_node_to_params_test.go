//ff:func feature=scan type=test topic=fastify control=iteration dimension=1
//ff:what schemaNodeToParams TypeBox ref 분기 → Param 변환 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestSchemaNodeToParams(t *testing.T) {
	fi := mustParse(t, []byte(`Type.Object({ page: Type.Integer() }); Query;`))
	innerObj := findAllByType(fi.Root, "object")[0]
	var queryID *sitter.Node
	for _, id := range findAllByType(fi.Root, "identifier") {
		if nodeText(id, fi.Src) == "Query" {
			queryID = id
		}
	}
	if queryID == nil {
		t.Fatal("no Query identifier")
	}
	vars := map[string]*sitter.Node{"Query": innerObj}
	params := schemaNodeToParams(queryID, fi.Src, vars)
	if len(params) != 1 || params[0].Name != "page" {
		t.Errorf("ref branch: %+v", params)
	}
}
