//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestCollectChainFromMemberExpr_Round5 테스트
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestCollectChainFromMemberExpr_Round5(t *testing.T) {

	root, src := parseTS(t, "const s = z.string;")
	var member *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if member == nil && n.Type() == "member_expression" {
			member = n
		}
	})
	if member == nil {
		t.Fatal("no member_expression")
	}
	var methods []ChainMethod
	collectChainFromMemberExpr(member, src, &methods)
	if len(methods) != 1 || methods[0].Name != "string" {
		t.Fatalf("expected [string], got %+v", methods)
	}
}
