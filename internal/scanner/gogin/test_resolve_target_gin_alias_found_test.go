//ff:func feature=scan type=test control=sequence
//ff:what TestResolveTargetGinAlias_Found 테스트
package gogin

import (
	"go/ast"
	rtgaPars "go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

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
