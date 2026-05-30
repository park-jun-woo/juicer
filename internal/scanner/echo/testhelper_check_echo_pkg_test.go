//ff:func feature=scan type=test control=sequence topic=echo
//ff:what checkEchoPkg 테스트 헬퍼
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

// checkEchoPkg type-checks an in-memory package at the echo import path so that
// named types resolve with the echo package suffix.
func checkEchoPkg(t *testing.T, src string) (*ast.File, *types.Info, *types.Package) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "echo.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	pkg, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	return file, info, pkg
}
