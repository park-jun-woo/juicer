//ff:func feature=scan type=test control=sequence topic=express
//ff:what isRouterStandaloneCall: 빈alias / Router() / new Router() / new Router / 미스매치
package express

import "testing"

func TestIsRouterStandaloneCall_EmptyAliases(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Router();`))
	if isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{}) {
		t.Fatal("expected false for empty aliases")
	}
}

func TestIsRouterStandaloneCall_PlainCall(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Router();`))
	if !isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected true for Router()")
	}
}

func TestIsRouterStandaloneCall_NewCall(t *testing.T) {
	// new Router() -> new_expression wrapping call_expression
	fi := mustParse(t, []byte(`const r = new Router();`))
	if !isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected true for new Router()")
	}
}

func TestIsRouterStandaloneCall_NewNoParens(t *testing.T) {
	// new Router (no parens) -> new_expression with identifier only -> target==node true
	fi := mustParse(t, []byte(`const r = new Router;`))
	v := findInitValue(firstDeclarator(t, fi))
	if v == nil {
		// some grammars still produce new_expression; locate it directly
		ne := findAllByType(fi.Root, "new_expression")
		if len(ne) == 0 {
			t.Skip("no new_expression produced")
		}
		v = ne[0]
	}
	if !isRouterStandaloneCall(v, fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected true for new Router")
	}
}

func TestIsRouterStandaloneCall_NotCallNode(t *testing.T) {
	// pass a non-call, non-new node (identifier) -> target.Type() != call_expression
	fi := mustParse(t, []byte(`Router;`))
	ids := findAllByType(fi.Root, "identifier")
	if isRouterStandaloneCall(ids[0], fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for identifier node")
	}
}

func TestIsRouterStandaloneCall_MemberCallee(t *testing.T) {
	// call with member_expression callee -> no direct identifier function child
	fi := mustParse(t, []byte(`a.b();`))
	if isRouterStandaloneCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for member callee")
	}
}

func TestIsRouterStandaloneCall_Mismatch(t *testing.T) {
	fi := mustParse(t, []byte(`const r = Other();`))
	if isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for Other()")
	}
}

func TestIsRouterStandaloneCall_NewMismatch(t *testing.T) {
	// new Other() -> unwrap returns call_expression Other() -> fn not in aliases
	fi := mustParse(t, []byte(`const r = new Other();`))
	if isRouterStandaloneCall(findInitValue(firstDeclarator(t, fi)), fi.Src, map[string]bool{"Router": true}) {
		t.Fatal("expected false for new Other()")
	}
}
