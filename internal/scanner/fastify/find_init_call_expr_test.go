//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what findInitCallExpr 테스트
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

func TestFindInitCallExpr_New(t *testing.T) {
	fi := mustParse(t, []byte("const x = new Foo();\n"))
	d := findAllByType(fi.Root, "variable_declarator")[0]
	c := findInitCallExpr(d)
	if c == nil || c.Type() != "new_expression" {
		t.Fatalf("expected new_expression, got %v", c)
	}
}

func TestFindInitCallExpr_None(t *testing.T) {
	fi := mustParse(t, []byte("const x = 5;\n"))
	d := findAllByType(fi.Root, "variable_declarator")[0]
	if c := findInitCallExpr(d); c != nil {
		t.Fatalf("expected nil for literal init, got %s", c.Type())
	}
}
