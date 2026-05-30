//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestInvokeIfMethodCall_WrongMethod 테스트
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestInvokeIfMethodCall_WrongMethod(t *testing.T) {
	src := []byte(`fn f() { x.route(a); }`)
	root, err := parseRust(src)
	if err != nil {
		t.Fatal(err)
	}
	call := findCallByFuncSuffix(root, src, ".route")
	fe := findChildByType(call, "field_expression")
	called := 0

	invokeIfMethodCall(call, fe, src, "service", func(args *sitter.Node) { called++ })
	if called != 0 {
		t.Fatalf("expected no callback, got %d", called)
	}
}
