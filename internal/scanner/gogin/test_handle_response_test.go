//ff:func feature=scan type=extract control=sequence
//ff:what TestHandleResponse 기본 응답 타입별 처리 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestHandleResponse(t *testing.T) {
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	t.Run("json response", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.INT, Value: "200"},
				&ast.Ident{Name: "result"},
			},
		}
		handleResponse(ep, call, "json", info, "handler")
		if len(ep.Responses) != 1 {
			t.Fatal("expected 1 response")
		}
		if ep.Responses[0].Status != "200" {
			t.Errorf("expected status 200, got %q", ep.Responses[0].Status)
		}
	})

	t.Run("string response", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.INT, Value: "200"},
			},
		}
		handleResponse(ep, call, "string", info, "handler")
		if ep.Responses[0].Kind != "string" {
			t.Errorf("expected kind 'string', got %q", ep.Responses[0].Kind)
		}
	})

	t.Run("data response", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{
			Args: []ast.Expr{
				&ast.BasicLit{Kind: token.INT, Value: "200"},
			},
		}
		handleResponse(ep, call, "data", info, "handler")
		if ep.Responses[0].Kind != "data" {
			t.Errorf("expected kind 'data'")
		}
	})

	t.Run("file response", func(t *testing.T) {
		ep := &scanner.Endpoint{}
		call := &ast.CallExpr{}
		handleResponse(ep, call, "file", info, "handler")
		if ep.Responses[0].Status != "200" {
			t.Errorf("expected status 200 for file")
		}
	})
}
