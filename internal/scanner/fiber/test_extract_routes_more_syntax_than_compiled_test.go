//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_MoreSyntaxThanCompiled 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestExtractRoutes_MoreSyntaxThanCompiled(t *testing.T) {

	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) { app.Get("/x", h) }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: nil,
		Fset:            fset,
	}
	eps, _ := extractRoutes([]*packages.Package{pkg}, "/proj")
	if len(eps) != 0 {
		t.Fatalf("expected 0 endpoints when no CompiledGoFiles, got %d", len(eps))
	}
}
