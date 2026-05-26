//ff:func feature=scan type=extract control=sequence
//ff:what TestIsGinRouterType_NotSelector 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestIsGinRouterType_NotSelector(t *testing.T) {
	if isGinRouterType(&ast.Ident{Name: "x"}, "gin") {
		t.Fatal("expected false for non-selector")
	}
}
