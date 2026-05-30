//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFindInitCallExpr_Call 테스트
package fastify

import "testing"

func TestFindInitCallExpr_Call(t *testing.T) {
	fi := mustParse(t, []byte("const app = Fastify();\n"))
	d := findAllByType(fi.Root, "variable_declarator")[0]
	c := findInitCallExpr(d)
	if c == nil || c.Type() != "call_expression" {
		t.Fatalf("expected call_expression, got %v", c)
	}
}
