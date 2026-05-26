//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_UnknownCov 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveStatusCode_UnknownCov(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	got := resolveStatusCode(expr, info)
	if got != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", got)
	}
}
