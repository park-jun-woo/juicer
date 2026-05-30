//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFindInitCallExpr_New 테스트
package fastify

import "testing"

func TestFindInitCallExpr_New(t *testing.T) {
	fi := mustParse(t, []byte("const x = new Foo();\n"))
	d := findAllByType(fi.Root, "variable_declarator")[0]
	c := findInitCallExpr(d)
	if c == nil || c.Type() != "new_expression" {
		t.Fatalf("expected new_expression, got %v", c)
	}
}
