//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextType_NotIdentCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_NotIdentCov(t *testing.T) {
	expr := &ast.StarExpr{X: &ast.SelectorExpr{X: &ast.CompositeLit{}, Sel: &ast.Ident{Name: "Context"}}}
	if isGinContextType(expr) {
		t.Fatal("expected false")
	}
}
