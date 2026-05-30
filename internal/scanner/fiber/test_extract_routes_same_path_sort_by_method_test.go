//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_SamePathSortByMethod 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestExtractRoutes_SamePathSortByMethod(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) {
	app.Post("/users", handler)
	app.Get("/users", handler)
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "/proj/m.go", src, 0)
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"/proj/m.go"},
		Fset:            fset,
	}
	eps, _ := extractRoutes([]*packages.Package{pkg}, "/proj")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}

	if eps[0].Method != "GET" || eps[1].Method != "POST" {
		t.Fatalf("not sorted by method: %s, %s", eps[0].Method, eps[1].Method)
	}
}
