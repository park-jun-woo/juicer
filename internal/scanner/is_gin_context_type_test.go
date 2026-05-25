//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinContextType_Valid 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_Valid(t *testing.T) {
	expr := &ast.StarExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "gin"},
			Sel: &ast.Ident{Name: "Context"},
		},
	}
	if !isGinContextType(expr) {
		t.Fatal("expected true")
	}
}
