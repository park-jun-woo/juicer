//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestRegisterParams_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestRegisterParams_Round5(t *testing.T) {
	src := `package m
import "github.com/labstack/echo/v4"
func setup(e *echo.Echo, g *echo.Group) {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var fn *ast.FuncDecl
	for _, d := range file.Decls {
		if f, ok := d.(*ast.FuncDecl); ok {
			fn = f
		}
	}
	routers := map[string]*routerInfo{}
	registerParams(fn, "echo", routers)
	if _, ok := routers["e"]; !ok {
		t.Fatalf("e not registered: %v", routers)
	}
	if _, ok := routers["g"]; !ok {
		t.Fatalf("g not registered: %v", routers)
	}
}
