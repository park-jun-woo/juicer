//ff:func feature=scan type=test topic=fastify control=iteration dimension=1
//ff:what schemaNodeToFields TypeBox ref/JSON 스키마 분기 → Field 변환 테스트
package fastify

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestSchemaNodeToFields(t *testing.T) {
	// Single source so all node byte offsets reference the same src.
	fi := mustParse(t, []byte(`Type.Object({ name: Type.String() }); Schema;`))
	innerObj := findAllByType(fi.Root, "object")[0]
	var schemaID *sitter.Node
	for _, id := range findAllByType(fi.Root, "identifier") {
		if nodeText(id, fi.Src) == "Schema" {
			schemaID = id
		}
	}
	if schemaID == nil {
		t.Fatal("no Schema identifier")
	}
	vars := map[string]*sitter.Node{"Schema": innerObj}
	fields := schemaNodeToFields(schemaID, fi.Src, vars)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Errorf("ref branch: %+v", fields)
	}
}
