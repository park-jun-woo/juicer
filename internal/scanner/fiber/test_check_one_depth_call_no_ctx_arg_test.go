//ff:func feature=scan type=test control=sequence
//ff:what TestCheckOneDepthCall_NoCtxArg 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"testing"
)

func TestCheckOneDepthCall_NoCtxArg(t *testing.T) {
	ep := &scanner.Endpoint{}

	call := parseCall(t, "doThing(x, y)")
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	checkOneDepthCall(ep, call, "c", newEmptyInfo(), idx)
	if len(ep.Responses) != 0 {
		t.Errorf("expected no responses, got %d", len(ep.Responses))
	}
}
