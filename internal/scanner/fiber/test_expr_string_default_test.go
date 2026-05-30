//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_Default 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestExprString_Default(t *testing.T) {

	got := exprString(&ast.Ellipsis{})
	if got == "" {
		t.Error("expected non-empty default representation")
	}
}
