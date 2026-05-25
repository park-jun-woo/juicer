package scanner

import (
	"go/ast"
	"go/token"
	"go/types"
	"testing"
)

func TestHandleResponse_JSON(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "200"},
			&ast.Ident{Name: "data"},
		},
	}
	info := &types.Info{
		Uses:  make(map[*ast.Ident]types.Object),
		Types: make(map[ast.Expr]types.TypeAndValue),
	}
	handleResponse(ep, call, "json", info, "handler")
	if len(ep.Responses) != 1 {
		t.Fatal("expected 1 response")
	}
}

func TestHandleResponse_File(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"file.txt"`}}}
	handleResponse(ep, call, "file", nil, "handler")
	if len(ep.Responses) != 1 || ep.Responses[0].Status != "200" {
		t.Fatal("expected 200 for file response")
	}
}

func TestHandleResponse_Status(t *testing.T) {
	ep := &Endpoint{}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "204"}}}
	handleResponse(ep, call, "status", nil, "handler")
	if len(ep.Responses) != 1 {
		t.Fatal("expected 1 response")
	}
}
