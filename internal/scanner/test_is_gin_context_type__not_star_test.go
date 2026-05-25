//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinContextType_NotStar 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinContextType_NotStar(t *testing.T) {
	if isGinContextType(&ast.Ident{Name: "int"}) {
		t.Fatal("expected false")
	}
}
