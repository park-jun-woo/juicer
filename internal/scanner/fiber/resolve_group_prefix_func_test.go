//ff:func feature=scan type=test control=iteration dimension=1
//ff:what resolveGroupPrefixFunc — 함수 단위 그룹 prefix 처리 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestResolveGroupPrefixFunc(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) {
	app.Get("/x", h)
	registerMore(app)
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", src, 0)
	var fn *ast.FuncDecl
	for _, d := range file.Decls {
		if f, ok := d.(*ast.FuncDecl); ok {
			fn = f
		}
	}
	pkg := &packages.Package{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, Fset: fset}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	ep := map[struct {
		file string
		line int
	}]int{}
	// should register "app" param and walk body without panic
	resolveGroupPrefixFunc(fn, "fiber", pkg, []*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{}, ep)
}
