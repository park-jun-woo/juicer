//ff:func feature=scan type=test control=sequence
//ff:what TestHandlerFuncName — 핸들러 표현에서 함수명 추출 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestHandlerFuncName(t *testing.T) {
	ident := &ast.Ident{Name: "GetUsers"}
	if got := handlerFuncName(ident); got != "GetUsers" {
		t.Errorf("ident: want GetUsers, got %s", got)
	}
	sel := &ast.SelectorExpr{X: &ast.Ident{Name: "h"}, Sel: &ast.Ident{Name: "List"}}
	if got := handlerFuncName(sel); got != "List" {
		t.Errorf("selector: want List, got %s", got)
	}
	call := &ast.CallExpr{Fun: ident}
	if got := handlerFuncName(call); got != "GetUsers" {
		t.Errorf("call: want GetUsers, got %s", got)
	}
	if got := handlerFuncName(&ast.BasicLit{Value: "1"}); got != "" {
		t.Errorf("basiclit: want empty, got %s", got)
	}
}
