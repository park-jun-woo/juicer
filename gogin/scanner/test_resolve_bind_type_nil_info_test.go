//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveBindType_NilInfo 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveBindType_NilInfo(t *testing.T) {
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.Ident{Name: "req"}},
	}
	typeName, fields := resolveBindType(call, nil)
	if typeName != "" || fields != nil {
		t.Error("expected empty for nil info")
	}
}
