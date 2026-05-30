//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_XNotIdent 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_XNotIdent(t *testing.T) {
	expr := &ast.SelectorExpr{X: &ast.CallExpr{}, Sel: &ast.Ident{Name: "App"}}
	if isFiberRouterType(expr, "fiber") {
		t.Fatal("expected false when X not ident")
	}
}
