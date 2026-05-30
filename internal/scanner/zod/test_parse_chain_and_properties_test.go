//ff:func feature=scan type=test control=iteration dimension=1 topic=zod
//ff:what TestParseChainAndProperties 테스트
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestParseChainAndProperties(t *testing.T) {
	root, src := parseTS(t, `const s = z.object({ name: z.string().min(1), age: z.number().int() });`)
	calls := findAllByType(root, "call_expression")
	var objNode *sitter.Node
	for _, c := range calls {
		if IsObjectCall(c, src) {
			args := findChildByType(c, "arguments")
			objNode = findChildByType(args, "object")
			break
		}
	}
	if objNode == nil {
		t.Fatal("no object node")
	}
	fields := ParseObjectProperties(objNode, src)
	if len(fields) != 2 {
		t.Fatalf("got %d fields", len(fields))
	}
	if fields[0].Name != "name" || fields[0].Type != "string" {
		t.Fatalf("name field: %+v", fields[0])
	}
	if fields[0].MinLength == nil || *fields[0].MinLength != 1 {
		t.Fatalf("min: %v", fields[0].MinLength)
	}
}
