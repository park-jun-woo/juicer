//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestAnalyzeHandlers_EmptyExprs_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"testing"
)

func TestAnalyzeHandlers_EmptyExprs_Round5(t *testing.T) {
	eps := []scanner.Endpoint{{File: "a.go", Line: 1}}

	analyzeHandlers(nil, eps, "/root", map[int][]ast.Expr{}, buildFuncIndex(nil))
}
