//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPkgName_NoImport 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestEchoPkgName_NoImport(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"net/http"`}},
		},
	}
	got := echoPkgName(file)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
