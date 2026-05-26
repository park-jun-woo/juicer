//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveCallerArgs_NilParams 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestResolveCallerArgs_NilParams(t *testing.T) {
	fnDecl := &ast.FuncDecl{
		Type: &ast.FuncType{Params: nil},
	}
	call := &ast.CallExpr{}
	status, typeName, fields, confidence := resolveCallerArgs(fnDecl, call, nil, nil)
	if status != "" || typeName != "" || fields != nil || confidence != "" {
		t.Error("expected empty results for nil params")
	}
}
