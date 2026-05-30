//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestEchoRouterParamAtIndex_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestEchoRouterParamAtIndex_Round5(t *testing.T) {
	src := `package echo
type Echo struct{}
func setup(prefix string, e *Echo) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "echo.go", src, 0)
	conf := types.Config{}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var fn *ast.FuncDecl
	for _, d := range file.Decls {
		if f, ok := d.(*ast.FuncDecl); ok && f.Name.Name == "setup" {
			fn = f
		}
	}

	if got := echoRouterParamAtIndex(fn, info, 1); got != "e" {
		t.Fatalf("idx1: %q", got)
	}

	if got := echoRouterParamAtIndex(fn, info, 0); got != "" {
		t.Fatalf("idx0: %q", got)
	}
}
