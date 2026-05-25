package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestFindInfoForExpr_NilPkgs(t *testing.T) {
	expr := &ast.Ident{Name: "x"}
	result := findInfoForExpr(expr, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}

func TestFindInfoForExpr_NilTypesCanon(t *testing.T) {
	// findInfoForExpr skips packages with nil TypesInfo
	expr := &ast.Ident{Name: "x"}
	result := findInfoForExpr(expr, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}

func TestFindInfoForExpr_NoMatchCanon(t *testing.T) {
	result := findInfoForExpr(&ast.Ident{Name: "x"}, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}

func TestFindInfoForExpr_NilInfoCanon(t *testing.T) {
	info := &types.Info{}
	_ = info
	result := findInfoForExpr(&ast.Ident{Name: "x"}, nil)
	if result != nil {
		t.Fatal("expected nil")
	}
}
