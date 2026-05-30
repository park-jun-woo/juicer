//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_Nil 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestResolveCallTarget_Nil(t *testing.T) {
	if pos := resolveCallTarget(&ast.CallExpr{Fun: &ast.Ident{Name: "x"}}, nil); pos.IsValid() {
		t.Fatalf("expected NoPos for nil info, got valid pos %v", pos)
	}
}
