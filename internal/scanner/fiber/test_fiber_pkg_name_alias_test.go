//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPkgName_Alias 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberPkgName_Alias(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{
				Name: &ast.Ident{Name: "fb"},
				Path: &ast.BasicLit{Value: `"github.com/gofiber/fiber/v2"`},
			},
		},
	}
	got := fiberPkgName(file)
	if got != "fb" {
		t.Fatalf("expected fb, got %s", got)
	}
}
