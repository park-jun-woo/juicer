//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_Ident 테스트
package gogin

import (
	"go/ast"
	"go/token"
	gtkn "go/token"
	gpars "go/parser"
	gimp "go/importer"
	gtyp "go/types"
	"go/types"
	"testing"
)

func TestResolveCallTarget_Ident(t *testing.T) {
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "f"}}
	info := &types.Info{
		Uses:       make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}
	pos := resolveCallTarget(call, info)
	if pos != token.NoPos {
		t.Fatal("expected NoPos")
	}
}


func TestResolveCallTarget_IdentUseResolved(t *testing.T) {
	src := `package m
func target() {}
func caller() { target() }
`
	fset := gtkn.NewFileSet()
	file, _ := gpars.ParseFile(fset, "m.go", src, 0)
	conf := gtyp.Config{Importer: gimp.Default()}
	info := &gtyp.Info{Uses: map[*ast.Ident]gtyp.Object{}}
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
	if !resolveCallTarget(call, info).IsValid() {
		t.Fatal("expected valid pos for ident call")
	}
}

func TestResolveCallTarget_SelectorMethod(t *testing.T) {
	src := `package m
type T struct{}
func (t T) M() {}
func caller() { var t T; t.M() }
`
	fset := gtkn.NewFileSet()
	file, _ := gpars.ParseFile(fset, "m.go", src, 0)
	conf := gtyp.Config{Importer: gimp.Default()}
	info := &gtyp.Info{Uses: map[*ast.Ident]gtyp.Object{}, Selections: map[*ast.SelectorExpr]*gtyp.Selection{}}
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
	if !resolveCallTarget(call, info).IsValid() {
		t.Fatal("expected valid pos for method call")
	}
}
