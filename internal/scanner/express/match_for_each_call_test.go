//ff:func feature=scan type=test control=sequence topic=express
//ff:what matchForEachCall: 정상(arrow/function) + 각 미매칭 분기
package express

import "testing"

func TestMatchForEachCall_Arrow(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(r => { parent.use('/x', r); });`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "routes" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_FunctionExpr(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(function(r) { parent.use('/x', r); });`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "routes" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`forEach(cb);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_NoObjIdent(t *testing.T) {
	// object is a member_expression -> no direct identifier child
	fi := mustParse(t, []byte(`a.b.forEach(r => r);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_NoArgsNode(t *testing.T) {
	// tagged template -> no arguments node
	fi := mustParse(t, []byte("routes.forEach`x`;"))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_NotForEach(t *testing.T) {
	fi := mustParse(t, []byte(`routes.map(r => r);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, nil); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_NoArrowOrFunc(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(cbRef);`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "" {
		t.Fatalf("got %q", got)
	}
}

func TestMatchForEachCall_NoRouterUse(t *testing.T) {
	fi := mustParse(t, []byte(`routes.forEach(r => { doThing(r); });`))
	if got := matchForEachCall(firstCallExpr(t, fi), fi.Src, map[string]bool{"parent": true}); got != "" {
		t.Fatalf("got %q", got)
	}
}
