//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_NilInfo 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestCheckOneDepthCall_NilInfo(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := parseCall(t, "respond(c, 200)")
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}

	checkOneDepthCall(ep, call, "c", nil, idx)
	if len(ep.Responses) != 0 {
		t.Errorf("expected no responses, got %d", len(ep.Responses))
	}
}
