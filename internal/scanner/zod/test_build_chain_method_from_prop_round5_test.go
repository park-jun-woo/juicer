//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestBuildChainMethodFromProp_Round5 테스트
package zod

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildChainMethodFromProp_Round5(t *testing.T) {
	root, src := parseTS(t, "const s = z.string().min(3);")
	// find the .min(...) call: it's a call_expression whose function property is "min"
	var minCall *sitter.Node
	walkNodes(root, func(n *sitter.Node) {
		if n.Type() != "call_expression" {
			return
		}
		fn := resolveFunctionNode(n)
		if fn == nil || fn.Type() != "member_expression" {
			return
		}
		prop := fn.ChildByFieldName("property")
		if prop != nil && nodeText(prop, src) == "min" {
			minCall = n
		}
	})
	if minCall == nil {
		t.Fatal("no min call")
	}
	fn := resolveFunctionNode(minCall)
	prop := fn.ChildByFieldName("property")
	cm := buildChainMethodFromProp(minCall, prop, src)
	if cm.Name != "min" {
		t.Fatalf("name: got %q", cm.Name)
	}
	if len(cm.Args) != 1 || cm.Args[0] != "3" {
		t.Fatalf("args: got %+v", cm.Args)
	}
}
