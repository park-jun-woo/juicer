//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprConst_Nil 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestResolveExprConst_Nil(t *testing.T) {
	if got := resolveExprConst(nil, &ast.Ident{Name: "x"}); got != "" {
		t.Fatalf("expected empty for nil info, got %s", got)
	}
}
