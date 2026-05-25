//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinHSelector_NonSelector 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinHSelector_NonSelector(t *testing.T) {
	if isGinHSelector(&ast.Ident{Name: "x"}) {
		t.Fatal("expected false")
	}
}
