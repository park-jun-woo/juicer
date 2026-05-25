//ff:func feature=scan type=extract control=sequence
//ff:what TestHandlePathParam_Duplicate 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandlePathParam_Duplicate(t *testing.T) {
	ep := &Endpoint{Request: &Request{PathParams: []Param{{Name: "id", Type: "string"}}}}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"id"`}},
	}
	handlePathParam(ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatal("should not duplicate")
	}
}
