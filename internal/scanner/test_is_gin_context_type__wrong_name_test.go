//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinContextType_WrongName 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_WrongName(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "http"},
			Sel: &ast.Ident{Name: "Request"},
		},
	}
	if isGinContextType(expr) {
		t.Fatal("expected false")
	}
}
