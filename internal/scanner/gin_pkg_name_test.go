//ff:func feature=scan type=test control=sequence
//ff:what TestGinPkgName_DefaultImport 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinPkgName_DefaultImport(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"github.com/gin-gonic/gin"`}},
		},
	}
	got := ginPkgName(file)
	if got != "gin" {
		t.Fatalf("expected gin, got %s", got)
	}
}

