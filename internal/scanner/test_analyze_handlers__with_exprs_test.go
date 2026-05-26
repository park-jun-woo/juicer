//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_WithExprs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestAnalyzeHandlers_WithExprs(t *testing.T) {
	eps := []Endpoint{{
		Path:   "/test",
		Method: "GET",
	}}
	handlerExprsMap := map[int][]ast.Expr{
		0: {&ast.Ident{Name: "handler"}},
	}
	analyzeHandlers(nil, eps, ".", handlerExprsMap)
}
