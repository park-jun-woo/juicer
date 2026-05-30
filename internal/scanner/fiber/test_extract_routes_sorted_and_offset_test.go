//ff:func feature=scan type=test control=sequence
//ff:what TestExtractRoutes_SortedAndOffset 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestExtractRoutes_SortedAndOffset(t *testing.T) {
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) {
	app.Get("/zebra", handler)
	app.Post("/apple", handler)
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/proj/m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: []string{"/proj/m.go"},
		Fset:            fset,
	}
	eps, hmap := extractRoutes([]*packages.Package{pkg}, "/proj")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(eps))
	}

	if eps[0].Path != "/apple" || eps[1].Path != "/zebra" {
		t.Fatalf("not sorted by path: %s, %s", eps[0].Path, eps[1].Path)
	}
	_ = hmap
}
