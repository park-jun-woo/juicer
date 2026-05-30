//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetFiberAlias_FileNotFound 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveTargetFiberAlias_FileNotFound(t *testing.T) {
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", "package m\n", 0)
	ctx := &groupArgCtx{pkgs: []*packages.Package{{Syntax: []*ast.File{file}}}}
	if got := resolveTargetFiberAlias(file.End()+1000, ctx); got != "" {
		t.Fatalf("expected empty for unknown pos, got %q", got)
	}
}
