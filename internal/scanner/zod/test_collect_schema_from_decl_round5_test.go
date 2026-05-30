//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectSchemaFromDecl_Round5 테스트
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectSchemaFromDecl_Round5(t *testing.T) {
	root, src := parseTS(t, "const UserSchema = z.object({});")
	var decl *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if decl == nil && (n.Type() == "lexical_declaration" || n.Type() == "variable_declaration") {
			decl = n
		}
	})
	if decl == nil {
		t.Fatal("no declaration")
	}
	schemas := map[string]*sitter.Node{}
	CollectSchemaFromDecl(decl, src, schemas)
	if _, ok := schemas["UserSchema"]; !ok {
		t.Fatalf("UserSchema not collected: %+v", schemas)
	}
}
