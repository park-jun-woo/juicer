//ff:func feature=scan type=test control=sequence topic=express
//ff:what extractOneRoute: 정상추출 + 각 nil 분기 검증
package express

import "testing"

func TestExtractOneRoute_Valid(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', h);`))
	ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true})
	if ri == nil || ri.Method != "GET" || ri.Path != "/x" || ri.Router != "r" {
		t.Fatalf("got %+v", ri)
	}
}

func TestExtractOneRoute_NoMember(t *testing.T) {
	fi := mustParse(t, []byte(`foo('/x');`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil, got %+v", ri)
	}
}

func TestExtractOneRoute_UnregisteredRouter(t *testing.T) {
	fi := mustParse(t, []byte(`other.get('/x', h);`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for unregistered router, got %+v", ri)
	}
}

func TestExtractOneRoute_NotHttpMethod(t *testing.T) {
	fi := mustParse(t, []byte(`r.foo('/x', h);`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for non-http method, got %+v", ri)
	}
}

func TestExtractOneRoute_NoIdentifierObject(t *testing.T) {
	// member object is itself a member_expression -> no direct identifier
	fi := mustParse(t, []byte(`a.b.get('/x', h);`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"a": true, "b": true}); ri != nil {
		t.Fatalf("expected nil, got %+v", ri)
	}
}

func TestExtractOneRoute_NoArgumentsNode(t *testing.T) {
	// tagged template -> call_expression without arguments node
	fi := mustParse(t, []byte("r.get`tpl`;"))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for no arguments node, got %+v", ri)
	}
}

func TestExtractOneRoute_NoArgsBuildNil(t *testing.T) {
	// registered router + http method but empty args -> buildRouteFromArgs nil
	fi := mustParse(t, []byte(`r.get();`))
	if ri := extractOneRoute(firstCallExpr(t, fi), fi.Src, map[string]bool{"r": true}); ri != nil {
		t.Fatalf("expected nil for no args, got %+v", ri)
	}
}
