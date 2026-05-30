//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestFindInitCallExpr_None 테스트
package fastify

import "testing"

func TestFindInitCallExpr_None(t *testing.T) {
	fi := mustParse(t, []byte("const x = 5;\n"))
	d := findAllByType(fi.Root, "variable_declarator")[0]
	if c := findInitCallExpr(d); c != nil {
		t.Fatalf("expected nil for literal init, got %s", c.Type())
	}
}
