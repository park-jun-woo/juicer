//ff:func feature=scan type=test control=sequence
//ff:what resolveTargetFiberAlias — fiber alias 해석 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestResolveTargetFiberAlias_Found(t *testing.T) {
	src := `package m
import fb "github.com/gofiber/fiber/v2"
func F(app *fb.App) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	ctx := &groupArgCtx{pkgs: []*packages.Package{{Syntax: []*ast.File{file}}}}
	if got := resolveTargetFiberAlias(file.Pos()+1, ctx); got != "fb" {
		t.Fatalf("alias = %q, want fb", got)
	}
}

func TestResolveTargetFiberAlias_FileNotFound(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", "package m\n", 0)
	ctx := &groupArgCtx{pkgs: []*packages.Package{{Syntax: []*ast.File{file}}}}
	if got := resolveTargetFiberAlias(file.End()+1000, ctx); got != "" {
		t.Fatalf("expected empty for unknown pos, got %q", got)
	}
}
