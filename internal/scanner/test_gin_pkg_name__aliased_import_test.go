//ff:func feature=scan type=extract control=sequence
//ff:what TestGinPkgName_AliasedImport 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinPkgName_AliasedImport(t *testing.T) {
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
