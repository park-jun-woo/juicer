//ff:func feature=scan type=test control=iteration dimension=1
//ff:what resolveGroupPrefix — 그룹 prefix 전파 진입점 테스트
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

func TestResolveGroupPrefix(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/proj/m.go", "package m\nfunc F() {}\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	pkgs := []*packages.Package{
		// nil TypesInfo -> skipped
		{Syntax: []*ast.File{file}, TypesInfo: nil, CompiledGoFiles: []string{"/proj/m.go"}, Fset: fset},
		// Syntax longer than CompiledGoFiles -> i>=len continue
		{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, CompiledGoFiles: nil, Fset: fset},
	}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	// should not panic
	resolveGroupPrefix(pkgs, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{})
}

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
