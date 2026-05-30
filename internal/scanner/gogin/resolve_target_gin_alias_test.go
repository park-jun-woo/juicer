//ff:func feature=scan type=test control=sequence
//ff:what resolveTargetGinAlias 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	rtgaPars "go/parser"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestResolveTargetGinAlias(t *testing.T) {
	fset := token.NewFileSet()
	ctx := &groupArgCtx{
		fset: fset,
		pkgs: []*packages.Package{},
	}

	// no matching file -> empty
	got := resolveTargetGinAlias(token.Pos(1), ctx)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestResolveTargetGinAlias_Found(t *testing.T) {
	src := `package m
import g "github.com/gin-gonic/gin"
func F(r *g.Engine) {}
`
	fset := token.NewFileSet()
	file, _ := rtgaPars.ParseFile(fset, "m.go", src, 0)
	ctx := &groupArgCtx{pkgs: []*packages.Package{{Syntax: []*ast.File{file}}}}
	if got := resolveTargetGinAlias(file.Pos()+1, ctx); got != "g" {
		t.Fatalf("alias = %q, want g", got)
	}
}
