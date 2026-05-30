//ff:func feature=scan type=test control=sequence topic=actix
//ff:what invokeIfMethodCall — 대상 메서드 호출 시에만 콜백 적용을 검증
package actix

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func TestInvokeIfMethodCall_Match(t *testing.T) {
	src := []byte(`fn f() { x.route(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	if call == nil {
		t.Fatal("no .route call")
	}
	fe := findChildByType(call, "field_expression")
	called := 0
	invokeIfMethodCall(call, fe, src, "route", func(args *sitter.Node) { called++ })
	if called != 1 {
		t.Fatalf("expected callback once, got %d", called)
	}
}

func TestInvokeIfMethodCall_WrongMethod(t *testing.T) {
	src := []byte(`fn f() { x.route(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	fe := findChildByType(call, "field_expression")
	called := 0
	// Asking for "service" against a ".route" call -> no callback.
	invokeIfMethodCall(call, fe, src, "service", func(args *sitter.Node) { called++ })
	if called != 0 {
		t.Fatalf("expected no callback, got %d", called)
	}
}

func TestInvokeIfMethodCall_NoArguments(t *testing.T) {
	// Build a method-call field_expression but pass a node (the field_expression
	// itself) that has no "arguments" child for the args lookup.
	src := []byte(`fn f() { x.route(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	fe := findChildByType(call, "field_expression")
	called := 0
	// Pass fe as the "n" node: it matches method but has no arguments child.
	invokeIfMethodCall(fe, fe, src, "route", func(args *sitter.Node) { called++ })
	if called != 0 {
		t.Fatalf("expected no callback when no arguments node, got %d", called)
	}
}
