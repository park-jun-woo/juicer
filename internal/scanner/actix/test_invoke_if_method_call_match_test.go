//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestInvokeIfMethodCall_Match 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
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
