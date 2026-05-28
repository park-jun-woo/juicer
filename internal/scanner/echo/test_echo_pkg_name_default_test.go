//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPkgName_Default 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestEchoPkgName_Default(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"github.com/labstack/echo/v4"`}},
		},
	}
	got := echoPkgName(file)
	if got != "echo" {
		t.Fatalf("expected echo, got %s", got)
	}
}
