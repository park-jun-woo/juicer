//ff:func feature=scan type=test control=selection
//ff:what resolveCallTarget — 호출 대상 위치 해석 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallTarget_IdentUse(t *testing.T) {
	src := `package m
func target() {}
func caller() { target() }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Uses: map[*ast.Ident]types.Object{}}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			call = c
		}
		return true
	})
	pos := resolveCallTarget(call, info)
	if !pos.IsValid() {
		t.Fatal("expected valid target pos for ident call")
	}
}

func TestResolveCallTarget_SelectorMethod(t *testing.T) {
	src := `package m
type T struct{}
func (t T) M() {}
func caller() { var t T; t.M() }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if _, ok := c.Fun.(*ast.SelectorExpr); ok {
				call = c
			}
		}
		return true
	})
	if call == nil {
		t.Fatal("method call not found")
	}
	if pos := resolveCallTarget(call, info); !pos.IsValid() {
		t.Fatal("expected valid pos for method call")
	}
}

func TestResolveCallTarget_SelectorPackageFunc(t *testing.T) {
	src := `package m
import "strings"
func caller() { strings.ToUpper("x") }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Uses:       map[*ast.Ident]types.Object{},
		Selections: map[*ast.SelectorExpr]*types.Selection{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if _, ok := c.Fun.(*ast.SelectorExpr); ok {
				call = c
			}
		}
		return true
	})
	// strings.ToUpper -> resolves via info.Uses[sel] path
	if pos := resolveCallTarget(call, info); !pos.IsValid() {
		t.Fatal("expected valid pos for package func call")
	}
}

func TestResolveCallTarget_Unresolved(t *testing.T) {
	// ident not in info.Uses -> NoPos
	call := parseCall(t, "unknown()")
	if pos := resolveCallTarget(call, newEmptyInfo()); pos.IsValid() {
		t.Fatal("expected NoPos for unresolved ident")
	}
}

func TestResolveCallTarget_DefaultFun(t *testing.T) {
	// call whose Fun is neither Ident nor SelectorExpr (a func literal call)
	call := parseCall(t, "(func(){})()")
	if pos := resolveCallTarget(call, newEmptyInfo()); pos.IsValid() {
		t.Fatal("expected NoPos for non-name fun")
	}
}
