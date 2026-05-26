//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinHSelector_NotSelectorCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinHSelector_NotSelectorCov(t *testing.T) {
	if isGinHSelector(&ast.Ident{Name: "x"}) {
		t.Fatal("expected false")
	}
}
