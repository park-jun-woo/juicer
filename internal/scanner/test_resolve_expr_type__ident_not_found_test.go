//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveExprType_IdentNotFound 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveExprType_IdentNotFound(t *testing.T) {
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
		Defs: make(map[*ast.Ident]types.Object),
	}
	tn, fields := resolveExprType(&ast.Ident{Name: "x"}, info)
	if tn != "" || fields != nil {
		t.Fatal("expected empty")
	}
}
