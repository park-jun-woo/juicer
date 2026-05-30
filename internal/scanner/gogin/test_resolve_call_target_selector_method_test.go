//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_SelectorMethod 테스트
package gogin

import (
	"go/ast"
	gimp "go/importer"
	gpars "go/parser"
	gtkn "go/token"
	gtyp "go/types"
	"testing"
)

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
