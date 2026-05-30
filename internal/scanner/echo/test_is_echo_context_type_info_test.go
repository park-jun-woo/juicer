//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIsEchoContextTypeInfo 테스트
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestIsEchoContextTypeInfo(t *testing.T) {

	src := `package echo
type Context interface{ Foo() }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "echo.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, nil)
	if err != nil {
		t.Fatal(err)
	}
	obj := pkg.Scope().Lookup("Context")
	if !isEchoContextTypeInfo(obj.Type()) {
		t.Fatal("expected Context type to match")
	}
}
