//ff:func feature=scan type=test control=sequence
//ff:what TestGinPkgName_AliasedCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestGinPkgName_AliasedCov(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Name: &ast.Ident{Name: "g"}, Path: &ast.BasicLit{Value: `"github.com/gin-gonic/gin"`}},
		},
	}
	got := ginPkgName(file)
	if got != "g" {
		t.Fatalf("expected g, got %s", got)
	}
}
