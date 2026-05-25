package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestResolveUsesConst_NotFound(t *testing.T) {
	info := &types.Info{Uses: make(map[*ast.Ident]types.Object)}
	got := resolveUsesConst(info, &ast.Ident{Name: "x"})
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
