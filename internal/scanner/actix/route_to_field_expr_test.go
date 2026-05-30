//ff:func feature=scan type=test control=sequence topic=actix
//ff:what routeToFieldExpr — .to(...) 호출 판별 분기를 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestRouteToFieldExpr_To(t *testing.T) {
	src := []byte(`fn f() { web::get().to(h); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".to")
	if call == nil {
		t.Fatal("no .to call")
	}
	fe := routeToFieldExpr(call, src)
	if fe == nil || fe.Type() != "field_expression" {
		t.Fatalf("expected field_expression, got %v", fe)
	}
}

func TestRouteToFieldExpr_OtherMethod(t *testing.T) {
	// .route(...) call has a field_expression but identifier != "to".
	src := []byte(`fn f() { x.route(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	if fe := routeToFieldExpr(call, src); fe != nil {
		t.Fatalf("expected nil for non-to method, got %v", fe)
	}
}

func TestRouteToFieldExpr_NoFieldExpr(t *testing.T) {
	// web::get() call's function is a scoped_identifier, not a field_expression.
	src := []byte(`fn f() { web::get(); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, "::get")
	if fe := routeToFieldExpr(call, src); fe != nil {
		t.Fatalf("expected nil when no field_expression, got %v", fe)
	}
}

func TestRouteToFieldExpr_NotCall(t *testing.T) {
	// A plain identifier node is not a call_expression.
	src := []byte(`fn f() { lone; }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	var id *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if id != nil {
			return
		}
		if n.Type() == "identifier" && nodeText(n, src) == "lone" {
			id = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	if id == nil {
		t.Fatal("identifier not found")
	}
	if fe := routeToFieldExpr(id, src); fe != nil {
		t.Fatalf("expected nil for non-call node, got %v", fe)
	}
}
