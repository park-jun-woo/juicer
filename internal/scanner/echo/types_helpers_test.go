//ff:func feature=scan type=test control=sequence topic=echo
//ff:what echo AST/types 헬퍼 함수 테스트 (in-memory go/types)
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestBindVarName(t *testing.T) {
	if bindVarName(parseExpr(t, "&dto")) != "dto" {
		t.Fatal("address-of")
	}
	if bindVarName(parseExpr(t, "dto")) != "dto" {
		t.Fatal("plain")
	}
}

func TestCompositeLitTypeName_Variants(t *testing.T) {
	d, b := compositeLitTypeName("UserDto{}")
	if d != "UserDto" || b != "UserDto" {
		t.Fatalf("got %q %q", d, b)
	}
	d2, b2 := compositeLitTypeName("[]UserDto{}")
	if d2 != "[]UserDto" || b2 != "UserDto" {
		t.Fatalf("slice: %q %q", d2, b2)
	}
	if d3, _ := compositeLitTypeName("notcomposite"); d3 != "" {
		t.Fatalf("non-composite: %q", d3)
	}
	if d4, _ := compositeLitTypeName("pkg.Type{}"); d4 != "" {
		t.Fatalf("qualified should be rejected: %q", d4)
	}
}

func TestExprString(t *testing.T) {
	if exprString(parseExpr(t, "foo")) != "foo" {
		t.Fatal("ident")
	}
	if exprString(parseExpr(t, "pkg.Type")) != "pkg.Type" {
		t.Fatal("selector")
	}
	if exprString(parseExpr(t, "T{}")) != "T{}" {
		t.Fatal("composite")
	}
}

func TestEchoPkgName(t *testing.T) {
	src := `package m
import "github.com/labstack/echo/v4"
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	if got := echoPkgName(file); got != "echo" {
		t.Fatalf("got %q", got)
	}

	src2 := `package m
import e "github.com/labstack/echo/v4"
`
	file2, _ := parser.ParseFile(fset, "m2.go", src2, 0)
	if got := echoPkgName(file2); got != "e" {
		t.Fatalf("aliased: %q", got)
	}

	src3 := `package m
import "fmt"
`
	file3, _ := parser.ParseFile(fset, "m3.go", src3, 0)
	if got := echoPkgName(file3); got != "" {
		t.Fatalf("no echo import: %q", got)
	}
}

func TestResolveUsesConst(t *testing.T) {
	src := `package m
const StatusOK = 200
var _ = StatusOK
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: map[*ast.Ident]types.Object{},
		Uses: map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	// find the use of StatusOK
	var useIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "StatusOK" {
			if _, isUse := info.Uses[id]; isUse {
				useIdent = id
			}
		}
		return true
	})
	if useIdent == nil {
		t.Fatal("no use found")
	}
	if got := resolveUsesConst(info, useIdent); got != "200" {
		t.Fatalf("got %q", got)
	}
}

func TestIsMapStringAny(t *testing.T) {
	src := `package m
var M = map[string]any{"a": 1}
var N = []int{1, 2}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var mapLit, sliceLit *ast.CompositeLit
	ast.Inspect(file, func(n ast.Node) bool {
		if cl, ok := n.(*ast.CompositeLit); ok {
			if _, isMap := cl.Type.(*ast.MapType); isMap {
				mapLit = cl
			} else {
				sliceLit = cl
			}
		}
		return true
	})
	if mapLit == nil {
		t.Fatal("no map literal")
	}
	if !isMapStringAny(mapLit, info) {
		t.Fatal("expected map[string]any to match")
	}
	if sliceLit != nil && isMapStringAny(sliceLit, info) {
		t.Fatal("slice should not match")
	}
}

func TestIsEchoContextTypeInfo(t *testing.T) {
	// echo.Context is an interface; emulate a named type called Context
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
