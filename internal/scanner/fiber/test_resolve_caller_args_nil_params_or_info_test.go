//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArgs_NilParamsOrInfo 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveCallerArgs_NilParamsOrInfo(t *testing.T) {
	fn := &ast.FuncDecl{Type: &ast.FuncType{}}
	status, tn, f, _ := resolveCallerArgs(fn, &ast.CallExpr{}, newEmptyInfo(), newEmptyInfo())
	if status != "" || tn != "" || f != nil {
		t.Fatalf("nil params: %q %q %v", status, tn, f)
	}

	fn2 := &ast.FuncDecl{Type: &ast.FuncType{Params: &ast.FieldList{}}}
	if s, _, _, _ := resolveCallerArgs(fn2, &ast.CallExpr{}, newEmptyInfo(), nil); s != "" {
		t.Fatalf("nil calleeInfo should return empty")
	}
}
