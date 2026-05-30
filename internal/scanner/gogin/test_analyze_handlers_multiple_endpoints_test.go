//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_MultipleEndpoints 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestAnalyzeHandlers_MultipleEndpoints(t *testing.T) {

	eps := []scanner.Endpoint{
		{Method: "GET", Path: "/a"},
		{Method: "POST", Path: "/b"},
	}
	hmap := map[int][]ast.Expr{
		0: {&ast.Ident{Name: "ha"}},
	}
	analyzeHandlers(nil, eps, ".", hmap, nil)
	if eps[1].Request != nil {
		t.Fatal("endpoint without handler should be untouched")
	}
}
