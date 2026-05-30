//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterType_NotSelector 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIsFiberRouterType_NotSelector(t *testing.T) {
	if isFiberRouterType(&ast.Ident{Name: "x"}, "fiber") {
		t.Fatal("expected false for non-selector")
	}
}
