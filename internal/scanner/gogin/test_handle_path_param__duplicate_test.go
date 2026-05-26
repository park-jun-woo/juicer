//ff:func feature=scan type=extract control=sequence
//ff:what TestHandlePathParam_Duplicate 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandlePathParam_Duplicate(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{PathParams: []scanner.Param{{Name: "id", Type: "string"}}}}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"id"`}},
	}
	handlePathParam(ep, call)
	if len(ep.Request.PathParams) != 1 {
		t.Fatal("should not duplicate")
	}
}
