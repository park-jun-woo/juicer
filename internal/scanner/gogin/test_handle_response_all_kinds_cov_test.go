//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestHandleResponse_AllKindsCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse_AllKindsCov(t *testing.T) {
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	kinds := []struct{ kind string; args int }{
		{"string", 1}, {"data", 1}, {"file", 0}, {"redirect", 1}, {"status", 1},
	}
	for _, k := range kinds {
		ep := &scanner.Endpoint{}
		var args []ast.Expr
		for i := 0; i < k.args; i++ {
			args = append(args, &ast.BasicLit{Kind: token.INT, Value: "200"})
		}
		call := &ast.CallExpr{Args: args}
		handleResponse(ep, call, k.kind, info, "callee.Func")
		if len(ep.Responses) != 1 {
			t.Fatalf("expected 1 response for kind %s", k.kind)
		}
		if ep.Responses[0].Source != "callee.Func" {
			t.Fatalf("expected source for kind %s", k.kind)
		}
	}
}
