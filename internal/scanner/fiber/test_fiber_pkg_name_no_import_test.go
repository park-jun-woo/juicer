//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPkgName_NoImport 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberPkgName_NoImport(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"net/http"`}},
		},
	}
	got := fiberPkgName(file)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
