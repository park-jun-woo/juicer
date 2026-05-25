package scanner

import (
	"go/ast"
	"testing"
)

func TestAnalyzeHandlers_Empty(t *testing.T) {
	analyzeHandlers(nil, nil, ".")
	analyzeHandlers(nil, []Endpoint{}, ".")
}

func TestAnalyzeHandlers_NoHandlers(t *testing.T) {
	eps := []Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".")
}

func TestAnalyzeHandlers_WithExprs(t *testing.T) {
	eps := []Endpoint{{
		Path:         "/test",
		Method:       "GET",
		handlerExprs: []ast.Expr{&ast.Ident{Name: "handler"}},
	}}
	analyzeHandlers(nil, eps, ".")
	if eps[0].handlerExprs != nil {
		t.Fatal("expected handlerExprs to be nil after analysis")
	}
}
