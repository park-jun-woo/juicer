//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_AllEmpty 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestAnalyzeHandlers_AllEmpty(t *testing.T) {

	endpoints := []scanner.Endpoint{{Path: "/x"}}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, byName: map[string]*ast.FuncDecl{}}
	analyzeHandlers(nil, endpoints, "/root", map[int][]ast.Expr{}, idx)
	if endpoints[0].Request != nil || len(endpoints[0].Responses) != 0 {
		t.Errorf("expected untouched endpoint, got %+v", endpoints[0])
	}
}
