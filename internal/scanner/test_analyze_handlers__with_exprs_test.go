//ff:func feature=scan type=extract control=sequence
//ff:what TestAnalyzeHandlers_WithExprs 테스트
package scanner

import (
	"go/ast"
	"testing"
)

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
