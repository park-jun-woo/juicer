//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterType_NotSelectorCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType_NotSelectorCov(t *testing.T) {
	if isGinRouterType(&ast.Ident{Name: "x"}, "gin") {
		t.Fatal("expected false")
	}
}
