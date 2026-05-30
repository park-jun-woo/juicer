//ff:func feature=scan type=test control=sequence
//ff:what TestAnalyzeExpr_CallExprRecurses 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAnalyzeExpr_CallExprRecurses(t *testing.T) {

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

	analyzeExpr(ep, call, nil, idx)
}
