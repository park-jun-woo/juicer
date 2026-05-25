//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveStatusCode_NoMatch 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveStatusCode_NoMatch(t *testing.T) {
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	expr := &ast.Ident{Name: "x"}
	got := resolveStatusCode(expr, info)
	if got != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", got)
	}
}
