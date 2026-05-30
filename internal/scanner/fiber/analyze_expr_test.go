//ff:func feature=scan type=test control=iteration dimension=1
//ff:what analyzeExpr — 핸들러 표현 분석 분기 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func newEmptyInfo() *types.Info {
	return &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
}

// firstFuncLit parses a Go source and returns the first FuncLit found.
func firstFuncLit(t *testing.T, src string) *ast.FuncLit {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var lit *ast.FuncLit
	ast.Inspect(file, func(n ast.Node) bool {
		if fl, ok := n.(*ast.FuncLit); ok && lit == nil {
			lit = fl
			return false
		}
		return true
	})
	if lit == nil {
		t.Fatal("no FuncLit found")
	}
	return lit
}

func TestAnalyzeExpr_FuncLit(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
var _ = func(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"a": "b"})
}
`
	lit := firstFuncLit(t, src)
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	// info nil -> AST fallback resolves ctx name "c"; should not panic.
	analyzeExpr(ep, lit, nil, idx)
}

func TestAnalyzeExpr_FuncLitNoCtx(t *testing.T) {
	// FuncLit whose params do not include *fiber.Ctx -> ctxName "" -> early return
	src := `package m
var _ = func(x int) error { return nil }
`
	lit := firstFuncLit(t, src)
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	analyzeExpr(ep, lit, nil, idx)
	// no request/response should be populated
	if ep.Request != nil {
		t.Errorf("expected no request, got %v", ep.Request)
	}
}

func TestAnalyzeExpr_DefaultExpr(t *testing.T) {
	// a BasicLit matches none of the cases -> no-op, no panic
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	analyzeExpr(ep, &ast.BasicLit{Kind: token.INT, Value: "1"}, nil, idx)
	if ep.Request != nil || len(ep.Responses) != 0 {
		t.Errorf("default case should be a no-op, got %+v", ep)
	}
}

func TestAnalyzeExpr_CallExprRecurses(t *testing.T) {
	// A CallExpr whose Fun is a FuncLit: analyzeExpr recurses into e.Fun.
	src := `package m
import "github.com/gofiber/fiber/v2"
var _ = (func(c *fiber.Ctx) error { return nil })()
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok && call == nil {
			call = c
			return false
		}
		return true
	})
	if call == nil {
		t.Fatal("no CallExpr found")
	}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	// recurses into the inner FuncLit (e.Fun); info nil -> AST fallback, no panic.
	analyzeExpr(ep, call, nil, idx)
}

func TestAnalyzeExpr_SelectorNotInSelections(t *testing.T) {
	// A SelectorExpr not present in info.Selections -> early return.
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	sel := &ast.SelectorExpr{X: ast.NewIdent("h"), Sel: ast.NewIdent("Method")}
	analyzeExpr(ep, sel, newEmptyInfo(), idx)
	if ep.Request != nil {
		t.Errorf("expected no-op for unresolved selector, got %v", ep.Request)
	}
}

func TestAnalyzeExpr_IdentNilUses(t *testing.T) {
	// An *ast.Ident with nil info -> info.Uses access guarded; reaches the
	// Ident case and returns early when no use is found.
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}
	ep := &scanner.Endpoint{}
	// Use a non-nil but empty types.Info so info.Uses[e] is nil -> early return.
	info := newEmptyInfo()
	analyzeExpr(ep, ast.NewIdent("handler"), info, idx)
	if ep.Request != nil {
		t.Errorf("expected no-op for unresolved ident, got %v", ep.Request)
	}
}
