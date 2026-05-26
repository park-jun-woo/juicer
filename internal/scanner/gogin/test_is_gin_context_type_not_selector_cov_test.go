//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextType_NotSelectorCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_NotSelectorCov(t *testing.T) {
	if isGinContextType(&ast.StarExpr{X: &ast.Ident{Name: "x"}}) {
		t.Fatal("expected false")
	}
}
