//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleQuery_DefaultQuery 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleQuery_DefaultQuery(t *testing.T) {
	ep := &scanner.Endpoint{}
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
