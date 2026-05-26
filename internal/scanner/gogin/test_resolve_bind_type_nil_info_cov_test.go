//ff:func feature=scan type=test control=sequence
//ff:what TestResolveBindType_NilInfoCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestResolveBindType_NilInfoCov(t *testing.T) {
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "x"}}}
	name, _ := resolveBindType(call, nil)
	if name != "" {
		t.Fatal("expected empty")
	}
}
