//ff:func feature=scan type=test control=sequence
//ff:what TestScanFile_NoGinImport 테스트
package scanner

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestScanFile_NoGinImport(t *testing.T) {
	file := &ast.File{
		Name:    &ast.Ident{Name: "main"},
		Imports: []*ast.ImportSpec{{Path: &ast.BasicLit{Value: `"fmt"`}}},
	}
	result := scanFile(file, "test.go", token.NewFileSet())
	if len(result) != 0 {
		t.Fatal("expected empty for non-gin file")
	}
}
