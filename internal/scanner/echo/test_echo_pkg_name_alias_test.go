//ff:func feature=scan type=test control=sequence
//ff:what TestEchoPkgName_Alias 테스트
package echo

import (
	"go/ast"
	"testing"
)

func TestEchoPkgName_Alias(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"github.com/labstack/echo/v4"`}, Name: &ast.Ident{Name: "e"}},
		},
	}
	got := echoPkgName(file)
	if got != "e" {
		t.Fatalf("expected e, got %s", got)
	}
}
