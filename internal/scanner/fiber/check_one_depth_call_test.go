//ff:func feature=scan type=test control=iteration dimension=1
//ff:what checkOneDepthCall — 1-depth 호출 추적 분기 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func parseCall(t *testing.T, expr string) *ast.CallExpr {
	t.Helper()
	e, err := parser.ParseExpr(expr)
	if err != nil {
		t.Fatal(err)
	}
	call, ok := e.(*ast.CallExpr)
	if !ok {
		t.Fatalf("%q is not a call", expr)
	}
	return call
}

func TestCheckOneDepthCall_NilInfo(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := parseCall(t, "respond(c, 200)")
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	// info nil -> immediate return, no panic, no responses
	checkOneDepthCall(ep, call, "c", nil, idx)
	if len(ep.Responses) != 0 {
		t.Errorf("expected no responses, got %d", len(ep.Responses))
	}
}

func TestCheckOneDepthCall_NoCtxArg(t *testing.T) {
	ep := &scanner.Endpoint{}
	// call does not pass ctxName "c" -> hasCtx false -> return
	call := parseCall(t, "doThing(x, y)")
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	checkOneDepthCall(ep, call, "c", newEmptyInfo(), idx)
	if len(ep.Responses) != 0 {
		t.Errorf("expected no responses, got %d", len(ep.Responses))
	}
}

func TestCheckOneDepthCall_HasCtxButUnresolved(t *testing.T) {
	ep := &scanner.Endpoint{}
	// call passes ctxName "c" -> hasCtx true, but empty info means
	// resolveCallTarget returns an invalid position -> return.
	call := parseCall(t, "respond(c, 200)")
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	checkOneDepthCall(ep, call, "c", newEmptyInfo(), idx)
	if len(ep.Responses) != 0 {
		t.Errorf("expected no responses, got %d", len(ep.Responses))
	}
}
