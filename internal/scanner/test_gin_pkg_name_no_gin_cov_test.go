//ff:func feature=scan type=test control=sequence
//ff:what TestGinPkgName_NoGinCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestGinPkgName_NoGinCov(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"fmt"`}},
		},
	}
	got := ginPkgName(file)
	if got != "" {
		t.Fatalf("expected empty, got %s", got)
	}
}
