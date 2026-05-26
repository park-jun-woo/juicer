//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_DuplicateCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/juicer/internal/scanner"
)

func TestHandleQuery_DuplicateCov(t *testing.T) {
	ep := &scanner.Endpoint{Request: &scanner.Request{Query: []scanner.Param{{Name: "page", Type: "string"}}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"page"`}}}
	handleQuery(ep, call, "Query")
	if len(ep.Request.Query) != 1 {
		t.Fatal("expected 1 (deduped)")
	}
}
