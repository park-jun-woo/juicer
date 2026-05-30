//ff:func feature=scan type=test control=iteration dimension=1
//ff:what analyzeHandlers — 핸들러 목록 분석 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeHandlers_EmptyAndFallback(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
var _ = func(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"a": "b"})
}
`
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

	endpoints := []scanner.Endpoint{{Path: "/a"}, {Path: "/b"}}
	// endpoint 0 has a handler expr; endpoint 1 has none (continue branch)
	handlerExprs := map[int][]ast.Expr{0: {lit}}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, byName: map[string]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}

	// empty pkgs -> findInfoForExpr nil -> analyzeExprFallback
	analyzeHandlers([]*packages.Package{}, endpoints, "/root", handlerExprs, idx)
	// should not panic; endpoint 1 untouched
	if endpoints[1].Request != nil {
		t.Errorf("endpoint with no handler should be untouched")
	}
}

func TestAnalyzeHandlers_AllEmpty(t *testing.T) {
	// every endpoint has no handler exprs -> all continue, nothing analyzed
	endpoints := []scanner.Endpoint{{Path: "/x"}}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, byName: map[string]*ast.FuncDecl{}}
	analyzeHandlers(nil, endpoints, "/root", map[int][]ast.Expr{}, idx)
	if endpoints[0].Request != nil || len(endpoints[0].Responses) != 0 {
		t.Errorf("expected untouched endpoint, got %+v", endpoints[0])
	}
}
