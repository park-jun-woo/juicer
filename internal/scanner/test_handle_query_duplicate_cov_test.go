//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_DuplicateCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleQuery_DuplicateCov(t *testing.T) {
	ep := &Endpoint{Request: &Request{Query: []Param{{Name: "page", Type: "string"}}}}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"page"`}}}
	handleQuery(ep, call, "Query")
	if len(ep.Request.Query) != 1 {
		t.Fatal("expected 1 (deduped)")
	}
}
