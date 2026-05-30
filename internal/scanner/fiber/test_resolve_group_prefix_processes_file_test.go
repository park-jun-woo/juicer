//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefix_ProcessesFile 테스트
package fiber

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveGroupPrefix_ProcessesFile(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) { app.Get("/x", h) }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", src, 0)
	pkgs := []*packages.Package{
		{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, CompiledGoFiles: []string{"/proj/m.go"}, Fset: fset},
	}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	resolveGroupPrefix(pkgs, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{})
}
