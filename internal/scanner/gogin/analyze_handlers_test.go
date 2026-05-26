//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_Empty 테스트
package gogin

import (
	"go/ast"
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
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
