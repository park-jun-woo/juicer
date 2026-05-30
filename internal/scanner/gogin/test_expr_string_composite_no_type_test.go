//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_CompositeNoType 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExprString_CompositeNoType(t *testing.T) {
	if got := exprString(&ast.CompositeLit{}); got != "{}" {
		t.Errorf("nil-type composite = %q, want {}", got)
	}
}
