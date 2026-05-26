//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_NilInfoCase 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_NilInfoCase(t *testing.T) {
	tn, fields := resolveExprType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || fields != nil {
		t.Fatal("expected empty")
	}
}

