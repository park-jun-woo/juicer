//ff:func feature=scan type=extract control=sequence
//ff:what TestExtractPathString 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestExtractPathString(t *testing.T) {
	t.Run("string literal", func(t *testing.T) {
		expr := &ast.BasicLit{Kind: token.STRING, Value: `"/users"`}
		got, ok := extractPathString(expr)
		if !ok || got != "/users" {
			t.Errorf("expected '/users', got %q, ok=%v", got, ok)
		}
	})

	t.Run("non-string literal", func(t *testing.T) {
		expr := &ast.BasicLit{Kind: token.INT, Value: "42"}
		_, ok := extractPathString(expr)
		if ok {
			t.Error("expected not ok")
		}
	})

	t.Run("string concatenation", func(t *testing.T) {
		expr := &ast.BinaryExpr{
			Op: token.ADD,
			X:  &ast.BasicLit{Kind: token.STRING, Value: `"/api"`},
			Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/v1"`},
		}
		got, ok := extractPathString(expr)
		if !ok || got != "/api/v1" {
			t.Errorf("expected '/api/v1', got %q, ok=%v", got, ok)
		}
	})

	t.Run("non-string concat", func(t *testing.T) {
		expr := &ast.BinaryExpr{
			Op: token.ADD,
			X:  &ast.Ident{Name: "baseURL"},
			Y:  &ast.BasicLit{Kind: token.STRING, Value: `"/v1"`},
		}
		got, ok := extractPathString(expr)
		if !ok || got != "/v1" {
			t.Errorf("expected '/v1', got %q, ok=%v", got, ok)
		}
	})

	t.Run("unknown expr", func(t *testing.T) {
		expr := &ast.Ident{Name: "path"}
		_, ok := extractPathString(expr)
		if ok {
			t.Error("expected not ok")
		}
	})
}
