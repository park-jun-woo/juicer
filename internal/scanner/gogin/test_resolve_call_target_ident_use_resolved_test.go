//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_IdentUseResolved 테스트
package gogin

import (
	"go/ast"
	gimp "go/importer"
	gpars "go/parser"
	gtkn "go/token"
	gtyp "go/types"
	"testing"
)

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
