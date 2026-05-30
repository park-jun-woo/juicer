//ff:func feature=scan type=test control=sequence
//ff:what TestIsGinRouterTypeInfo_NonPointer 테스트
package gogin

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestIsGinRouterTypeInfo_NonPointer(t *testing.T) {
	basic := types.Typ[types.Int]
	if isGinRouterTypeInfo(basic) {
		t.Fatal("expected false for non-pointer type")
	}
}

func TestIsGinRouterTypeInfo_NamedNonGinPkg(t *testing.T) {
	// a local type named "Engine" (in ginRouterTypes set) but pkg is not gin
	src := `package m
type Engine struct{}
var P = &Engine{}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("m", fset, []*ast.File{file}, &types.Info{})
	if err != nil {
		t.Fatal(err)
	}
	typ := pkg.Scope().Lookup("P").Type()
	if isGinRouterTypeInfo(typ) {
		t.Fatal("local *Engine (non-gin pkg) should be false")
	}
}

