//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveBindType_NilInfoCase 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestResolveBindType_NilInfoCase(t *testing.T) {
	call := &ast.CallExpr{Args: []ast.Expr{&ast.Ident{Name: "req"}}}
	name, fields := resolveBindType(call, nil)
	if name != "" || fields != nil {
		t.Fatal("expected empty with nil info")
	}
}
