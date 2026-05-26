//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveCallerArgs_NilCalleeInfo 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestResolveCallerArgs_NilCalleeInfo(t *testing.T) {
	fnDecl := &ast.FuncDecl{
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: []*ast.Field{{
					Names: []*ast.Ident{{Name: "x"}},
					Type:  &ast.Ident{Name: "int"},
				}},
			},
		},
	}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "200"}},
	}
	status, typeName, fields, confidence := resolveCallerArgs(fnDecl, call, nil, nil)
	if status != "" || typeName != "" || fields != nil || confidence != "" {
		t.Error("expected empty results for nil calleeInfo")
	}
}
