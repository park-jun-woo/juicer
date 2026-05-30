//ff:func feature=scan type=test control=sequence topic=express
//ff:what unwrapChain: route체인 base / 비call / 비멤버 / obj비call 분기
package express

import "testing"

func TestUnwrapChain_Base(t *testing.T) {
	fi := mustParse(t, []byte(`router.route('/:id').get(h);`))
	path, rv, methods := unwrapChain(outermostCall(fi), fi.Src, map[string]bool{"router": true})
	if path != "/:id" || rv != "router" || len(methods) != 1 {
		t.Fatalf("path=%q rv=%q methods=%v", path, rv, methods)
	}
}

func TestUnwrapChain_Recursive(t *testing.T) {
	// 3-level chain: outermost .put -> obj .get (call, not route) -> recurse
	fi := mustParse(t, []byte(`router.route('/:id').get(getH).put(putH);`))
	path, rv, methods := unwrapChain(outermostCall(fi), fi.Src, map[string]bool{"router": true})
	if path != "/:id" || rv != "router" || len(methods) != 2 {
		t.Fatalf("path=%q rv=%q methods=%v", path, rv, methods)
	}
}

func TestUnwrapChain_NotCall(t *testing.T) {
	fi := mustParse(t, []byte(`router;`))
	ids := findAllByType(fi.Root, "identifier")
	if p, _, m := unwrapChain(ids[0], fi.Src, map[string]bool{"router": true}); p != "" || m != nil {
		t.Fatalf("expected empty")
	}
}

func TestUnwrapChain_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`foo();`))
	if p, _, m := unwrapChain(firstCallExpr(t, fi), fi.Src, nil); p != "" || m != nil {
		t.Fatalf("expected empty")
	}
}

func TestUnwrapChain_ObjNotCall(t *testing.T) {
	// router.get(...) -> object is identifier, not a call_expression
	fi := mustParse(t, []byte(`router.get('/x');`))
	if p, _, m := unwrapChain(firstCallExpr(t, fi), fi.Src, map[string]bool{"router": true}); p != "" || m != nil {
		t.Fatalf("expected empty, got p=%q m=%v", p, m)
	}
}
