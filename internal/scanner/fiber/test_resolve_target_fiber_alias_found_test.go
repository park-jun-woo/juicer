//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetFiberAlias_Found 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
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
