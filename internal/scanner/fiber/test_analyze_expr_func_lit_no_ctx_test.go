//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_FuncLitNoCtx 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestAnalyzeExpr_FuncLitNoCtx(t *testing.T) {

	src := `package m
var _ = func(x int) error { return nil }
`
	lit := firstFuncLit(t, src)
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	analyzeExpr(ep, lit, nil, idx)

	if ep.Request != nil {
		t.Errorf("expected no request, got %v", ep.Request)
	}
}
