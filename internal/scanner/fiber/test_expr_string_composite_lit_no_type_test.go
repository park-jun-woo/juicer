//ff:func feature=scan type=test control=sequence
//ff:what TestExprString_CompositeLitNoType 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestExprString_CompositeLitNoType(t *testing.T) {

	got := exprStringFor(t, "[]Book{{Title: \"x\"}}")
	if got != "[]Book{}" {
		t.Errorf("outer composite = %q", got)
	}

	cl := &ast.CompositeLit{}
	if s := exprString(cl); s != "{}" {
		t.Errorf("nil-type composite = %q, want {}", s)
	}
}
