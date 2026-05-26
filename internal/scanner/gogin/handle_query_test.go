//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_Basic 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleQuery_Basic(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"page"`}},
	}
	handleQuery(ep, call, "Query")
	if len(ep.Request.Query) != 1 {
		t.Fatal("expected 1 query param")
	}
}

