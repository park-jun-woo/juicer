//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_IdentNilUses 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestAnalyzeExpr_IdentNilUses(t *testing.T) {

	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}

	info := newEmptyInfo()
	analyzeExpr(ep, ast.NewIdent("handler"), info, idx)
	if ep.Request != nil {
		t.Errorf("expected no-op for unresolved ident, got %v", ep.Request)
	}
}
