//ff:func feature=scan type=test control=iteration dimension=1 topic=zod
//ff:what TestParseSchemaAndObjectArgs 테스트
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestParseSchemaAndObjectArgs(t *testing.T) {
	root, src := parseTS(t, `const s = z.object({ name: z.string() });`)
	calls := findAllByType(root, "call_expression")
	var objCall *sitter.Node
	for _, c := range calls {
		if IsObjectCall(c, src) {
			objCall = c
			break
		}
	}
	if objCall == nil {
		t.Fatal("no object call")
	}
	fields := ParseObjectArgs(objCall, src)
	if len(fields) != 1 || fields[0].Name != "name" {
		t.Fatalf("ParseObjectArgs: %+v", fields)
	}
	schemaFields := ParseSchema(objCall, src)
	if len(schemaFields) != 1 {
		t.Fatalf("ParseSchema: %+v", schemaFields)
	}
}
