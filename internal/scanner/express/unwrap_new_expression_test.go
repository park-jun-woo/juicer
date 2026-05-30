//ff:func feature=scan type=test control=sequence topic=express
//ff:what unwrapNewExpression: call내부 / identifier+alias / 미매칭 분기
package express

import "testing"

func TestUnwrapNewExpression_NewRouterIdent(t *testing.T) {
	// new Router() -> no inner call_expression child; constructor identifier matches alias
	fi := mustParse(t, []byte(`const r = new Router();`))
	ne := findAllByType(fi.Root, "new_expression")[0]
	got := unwrapNewExpression(ne, fi.Src, map[string]bool{"Router": true})
	if got == nil || got != ne {
		t.Fatalf("expected node itself, got %v", got)
	}
}

func TestUnwrapNewExpression_IdentifierAlias(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Router;`))
	ne := findAllByType(fi.Root, "new_expression")
	if len(ne) == 0 {
		t.Skip("no new_expression for bare new")
	}
	got := unwrapNewExpression(ne[0], fi.Src, map[string]bool{"Router": true})
	if got == nil {
		t.Fatal("expected non-nil for identifier alias")
	}
}

// The `findChildByType(node, "call_expression")` branch is unreachable for a
// real `new_expression`: tree-sitter never makes a call_expression a *direct*
// child (constructor is identifier/member/parenthesized_expression).

func TestUnwrapNewExpression_NoMatch(t *testing.T) {
	fi := mustParse(t, []byte(`const r = new Other;`))
	ne := findAllByType(fi.Root, "new_expression")
	if len(ne) == 0 {
		t.Skip("no new_expression")
	}
	if got := unwrapNewExpression(ne[0], fi.Src, map[string]bool{"Router": true}); got != nil {
		t.Fatalf("expected nil, got %v", got.Type())
	}
}
