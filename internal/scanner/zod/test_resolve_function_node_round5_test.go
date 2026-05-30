//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestResolveFunctionNode_Round5 테스트
package zod

import "testing"

func TestResolveFunctionNode_Round5(t *testing.T) {
	root, _ := parseTS(t, "const s = z.string();")
	call := findFirstCall(t, root)
	fn := resolveFunctionNode(call)
	if fn == nil {
		t.Fatal("expected function node")
	}
	if fn.Type() != "member_expression" {
		t.Fatalf("expected member_expression, got %s", fn.Type())
	}
}
