//ff:func feature=scan type=test control=sequence
//ff:what TestHandleQuery_DefaultQueryCov 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleQuery_DefaultQueryCov(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{
		&ast.BasicLit{Kind: token.STRING, Value: `"page"`},
		&ast.BasicLit{Kind: token.STRING, Value: `"1"`},
	}}
	handleQuery(ep, call, "DefaultQuery")
	if ep.Request.Query[0].Default != "1" {
		t.Fatalf("expected default 1, got %s", ep.Request.Query[0].Default)
	}
}
