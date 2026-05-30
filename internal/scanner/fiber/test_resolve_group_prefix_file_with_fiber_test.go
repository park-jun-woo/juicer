//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefixFile_WithFiber 테스트
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

func TestResolveGroupPrefixFile_WithFiber(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) { app.Get("/x", h) }
var X = 1
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", src, 0)
	pkg := &packages.Package{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, Fset: fset}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}

	resolveGroupPrefixFile(file, pkg, []*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndexEmpty())
}
