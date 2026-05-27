//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_WithExprs 테스트
package gogin

import (
	"go/ast"
	"testing"
	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeHandlers_WithExprs(t *testing.T) {
	eps := []scanner.Endpoint{{
		Path:   "/test",
		Method: "GET",
	}}
	handlerExprsMap := map[int][]ast.Expr{
		0: {&ast.Ident{Name: "handler"}},
	}
	analyzeHandlers(nil, eps, ".", handlerExprsMap, nil)
}
