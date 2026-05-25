//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleResponseEdgeCases redirect, status, 미확인 상태 등 edge case 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestHandleResponseEdgeCases(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	t.Run("redirect response", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.INT, Value: "301"},
			},
		}
		handleResponse(ep, call, "redirect", info, "handler")
		if ep.Responses[0].Status != "301" {
			t.Errorf("expected status 301, got %q", ep.Responses[0].Status)
		}
	})

	t.Run("status response", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.INT, Value: "204"},
			},
		}
		handleResponse(ep, call, "status", info, "handler")
		if ep.Responses[0].Status != "204" {
			t.Errorf("expected status 204")
		}
	})

	t.Run("no status code", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{}
		handleResponse(ep, call, "json", info, "handler")
		if ep.Responses[0].Status != "(unknown)" {
			t.Errorf("expected (unknown)")
		}
	})

	t.Run("non-handler source", func(t *testing.T) {
		ep := &Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.INT, Value: "200"},
			},
		}
		handleResponse(ep, call, "json", info, "helper")
		if ep.Responses[0].Source != "helper" {
			t.Errorf("expected source 'helper', got %q", ep.Responses[0].Source)
		}
	})
}
