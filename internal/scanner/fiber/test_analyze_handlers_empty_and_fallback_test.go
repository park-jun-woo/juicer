//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeHandlers_EmptyAndFallback 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
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

	handlerExprs := map[int][]ast.Expr{0: {lit}}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}, byName: map[string]*ast.FuncDecl{}, astStructs: map[string]*ast.StructType{}}

	analyzeHandlers([]*packages.Package{}, endpoints, "/root", handlerExprs, idx)

	if endpoints[1].Request != nil {
		t.Errorf("endpoint with no handler should be untouched")
	}
}
