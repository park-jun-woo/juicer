//ff:func feature=scan type=test control=iteration dimension=1
//ff:what resolveGroupPrefixFile — 파일 단위 그룹 prefix 처리 테스트
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

func epIndexEmpty() map[struct {
	file string
	line int
}]int {
	return map[struct {
		file string
		line int
	}]int{}
}

func TestResolveGroupPrefixFile_NoFiber(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", "package m\nfunc F() {}\n", 0)
	pkg := &packages.Package{Syntax: []*ast.File{file}, TypesInfo: &types.Info{}, Fset: fset}
	idx := &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}}
	// no fiber import -> early return, no panic
	resolveGroupPrefixFile(file, pkg, []*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndexEmpty())
}

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
	// has fiber import; iterates decls (Setup func + non-func var X skipped)
	resolveGroupPrefixFile(file, pkg, []*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndexEmpty())
}
