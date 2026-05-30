//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestEchoCtxParamNameInfo_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// echoCtxFuncType returns the *ast.FuncType of the first func decl whose first
// param is an echo.Context, plus the type info.
func TestEchoCtxParamNameInfo_Round5(t *testing.T) {
	src := `package echo
type Context interface{ JSON(int, interface{}) error }
func handler(c Context) error { return nil }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "echo.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var ft *ast.FuncType
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "handler" {
			ft = fn.Type
		}
	}
	if ft == nil {
		t.Fatal("no handler func")
	}
	if got := echoCtxParamNameInfo(ft, info); got != "c" {
		t.Fatalf("got %q", got)
	}

	_ = echoCtxParamNameInfo(ft, nil)
}
