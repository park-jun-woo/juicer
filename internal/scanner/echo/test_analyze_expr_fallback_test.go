//ff:func feature=scan type=test topic=echo control=sequence
//ff:what analyzeExprFallback FuncLit/명명 핸들러 폴백 분석 및 미해석 무시 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestAnalyzeExprFallback(t *testing.T) {
	src := `package p
import "github.com/labstack/echo/v4"
func Handler(c echo.Context) error {
	return c.JSON(200, map[string]any{"ok": true})
}
func register(e *echo.Echo) {
	e.GET("/x", func(c echo.Context) error {
		return c.JSON(200, map[string]any{"a": 1})
	})
}`
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := &funcIndex{
		byName:     map[string]*ast.FuncDecl{},
		byPos:      map[token.Pos]*ast.FuncDecl{},
		astStructs: map[string]*ast.StructType{},
	}
	indexFileDecls(f, nil, idx)

	// FuncLit branch
	var lit *ast.FuncLit
	ast.Inspect(f, func(n ast.Node) bool {
		if l, ok := n.(*ast.FuncLit); ok && lit == nil {
			lit = l
		}
		return true
	})
	if lit == nil {
		t.Fatal("no FuncLit")
	}
	ep := &scanner.Endpoint{}
	analyzeExprFallback(ep, lit, idx)
	if len(ep.Responses) == 0 {
		t.Errorf("FuncLit should yield a response")
	}

	// named-handler branch
	ep2 := &scanner.Endpoint{}
	analyzeExprFallback(ep2, ast.NewIdent("Handler"), idx)
	if len(ep2.Responses) == 0 {
		t.Errorf("named handler should yield a response")
	}

	// unresolvable name -> no-op
	ep3 := &scanner.Endpoint{}
	analyzeExprFallback(ep3, ast.NewIdent("Ghost"), idx)
	if len(ep3.Responses) != 0 {
		t.Errorf("ghost should be no-op: %+v", ep3.Responses)
	}
}
