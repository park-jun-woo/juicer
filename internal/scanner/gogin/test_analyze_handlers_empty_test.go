//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_Empty 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestAnalyzeHandlers_Empty(t *testing.T) {
	analyzeHandlers(nil, nil, ".", nil, nil)
	analyzeHandlers(nil, []scanner.Endpoint{}, ".", nil, nil)

	eps := []scanner.Endpoint{{Method: "GET", Path: "/test"}}
	analyzeHandlers(nil, eps, ".", map[int][]ast.Expr{}, nil)

	analyzeHandlers(nil, eps, ".", map[int][]ast.Expr{0: {&ast.Ident{Name: "x"}}}, nil)
}
