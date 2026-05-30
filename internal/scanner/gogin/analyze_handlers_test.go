//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_Empty 테스트
package gogin

import (
	"go/ast"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeHandlers_Empty(t *testing.T) {
	analyzeHandlers(nil, nil, ".", nil, nil)
	analyzeHandlers(nil, []scanner.Endpoint{}, ".", nil, nil)

	// endpoint with no exprs (continue branch)
	eps := []scanner.Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".", map[int][]ast.Expr{}, nil)

	// endpoint with exprs but no info (nil info continue)
	analyzeHandlers(nil, eps, ".", map[int][]ast.Expr{0: {&ast.Ident{Name: "x"}}}, nil)
}

func TestAnalyzeHandlers_MultipleEndpoints(t *testing.T) {
	// two endpoints, second has no handler exprs -> exercises both loop iterations
	eps := []scanner.Endpoint{
		{Method: "GET", Path: "/a"},
		{Method: "POST", Path: "/b"},
	}
	hmap := map[int][]ast.Expr{
		0: {&ast.Ident{Name: "ha"}}, // nil info -> continue
	}
	analyzeHandlers(nil, eps, ".", hmap, nil)
	if eps[1].Request != nil {
		t.Fatal("endpoint without handler should be untouched")
	}
}
