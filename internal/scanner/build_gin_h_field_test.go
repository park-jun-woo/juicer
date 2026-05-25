package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestBuildGinHField_NonKV(t *testing.T) {
	f := buildGinHField(&ast.Ident{Name: "x"}, &types.Info{})
	if f != nil {
		t.Fatal("expected nil")
	}
}
