//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArgs_EmptyParams 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveCallerArgs_EmptyParams(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	call := &ast.CallExpr{}
	status, tn, fields, conf := resolveCallerArgs(fn, call, nil, nil)
	if status != "" || tn != "" || fields != nil || conf != "" {
		t.Fatal("expected empty results")
	}
}

