//ff:func feature=scan type=test topic=fiber control=sequence
//ff:what scanFallbackFunc/analyzeExprFallback fiber.Ctx 핸들러 응답 폴백 분석 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestFallbackFunc(t *testing.T) {
	src := `package p
import "github.com/gofiber/fiber/v2"
func H(c *fiber.Ctx) error { return c.JSON(map[string]any{"x": 1}) }
func register(app *fiber.App) {
	app.Get("/x", func(c *fiber.Ctx) error { return c.JSON(map[string]any{"a": 1}) })
}`
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := newFiberIndex()
	indexFileDecls(f, nil, idx)

	// scanFallbackFunc on named handler
	h := idx.byName["H"]
	ep := &scanner.Endpoint{}
	scanFallbackFunc(ep, h.Type, h.Body, idx)
	if len(ep.Responses) == 0 {
		t.Error("handler should produce responses")
	}

	// analyzeExprFallback FuncLit branch
	var lit *ast.FuncLit
	ast.Inspect(f, func(n ast.Node) bool {
		if l, ok := n.(*ast.FuncLit); ok && lit == nil {
			lit = l
		}
		return true
	})
	ep2 := &scanner.Endpoint{}
	analyzeExprFallback(ep2, lit, idx)
	if len(ep2.Responses) == 0 {
		t.Error("FuncLit should produce responses")
	}

	// analyzeExprFallback named + unresolvable
	ep3 := &scanner.Endpoint{}
	analyzeExprFallback(ep3, ast.NewIdent("H"), idx)
	if len(ep3.Responses) == 0 {
		t.Error("named handler should produce responses")
	}
	ep4 := &scanner.Endpoint{}
	analyzeExprFallback(ep4, ast.NewIdent("Ghost"), idx)
	if len(ep4.Responses) != 0 {
		t.Error("ghost should be no-op")
	}
}
