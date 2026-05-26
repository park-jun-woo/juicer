//ff:func feature=scan type=test control=sequence
//ff:what TestHandlePathParam_DuplicateCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandlePathParam_DuplicateCov(t *testing.T) {
	ep := &Endpoint{Request: &Request{PathParams: []Param{{Name: "id", Type: "string"}}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"id"`}}}
	handlePathParam(ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatal("expected 1 (deduped)")
	}
}
