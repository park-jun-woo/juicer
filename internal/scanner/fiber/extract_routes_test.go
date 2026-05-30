//ff:func feature=scan type=test control=iteration dimension=1
//ff:what extractRoutes — 라우트 추출/정렬 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
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
	// sorted by Path: /apple before /zebra
	if eps[0].Path != "/apple" || eps[1].Path != "/zebra" {
		t.Fatalf("not sorted by path: %s, %s", eps[0].Path, eps[1].Path)
	}
	_ = hmap
}

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
	// same path -> sort by method: GET before POST
	if eps[0].Method != "GET" || eps[1].Method != "POST" {
		t.Fatalf("not sorted by method: %s, %s", eps[0].Method, eps[1].Method)
	}
}

func TestExtractRoutes_MoreSyntaxThanCompiled(t *testing.T) {
	// pkg with a Syntax file but empty CompiledGoFiles -> i>=len continue
	src := `package m
import "github.com/gofiber/fiber/v2"
func Setup(app *fiber.App) { app.Get("/x", h) }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		CompiledGoFiles: nil, // shorter than Syntax -> continue
		Fset:            fset,
	}
	eps, _ := extractRoutes([]*packages.Package{pkg}, "/proj")
	if len(eps) != 0 {
		t.Fatalf("expected 0 endpoints when no CompiledGoFiles, got %d", len(eps))
	}
}
