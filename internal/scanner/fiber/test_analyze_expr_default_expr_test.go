//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_DefaultExpr 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestAnalyzeExpr_DefaultExpr(t *testing.T) {

	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	analyzeExpr(ep, &ast.BasicLit{Kind: token.INT, Value: "1"}, nil, idx)
	if ep.Request != nil || len(ep.Responses) != 0 {
		t.Errorf("default case should be a no-op, got %+v", ep)
	}
}
