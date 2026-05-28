//ff:func feature=scan type=test control=sequence
//ff:what TestFiberPkgName_Default 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestFiberPkgName_Default(t *testing.T) {
	file := &ast.File{
		Imports: []*ast.ImportSpec{
			{Path: &ast.BasicLit{Value: `"github.com/gofiber/fiber/v2"`}},
		},
	}
	got := fiberPkgName(file)
	if got != "fiber" {
		t.Fatalf("expected fiber, got %s", got)
	}
}
