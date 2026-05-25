//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleQuery_DefaultQuery 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestHandleQuery_DefaultQuery(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: `"page"`},
			&ast.BasicLit{Kind: token.STRING, Value: `"1"`},
		},
	}
	handleQuery(ep, call, "DefaultQuery")
	if len(ep.Request.Query) != 1 || ep.Request.Query[0].Default != "1" {
		t.Fatal("expected default value")
	}
}
