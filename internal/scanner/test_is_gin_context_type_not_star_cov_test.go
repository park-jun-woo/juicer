//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinContextType_NotStarCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_NotStarCov(t *testing.T) {
	if isGinContextType(&ast.Ident{Name: "x"}) {
		t.Fatal("expected false")
	}
}
