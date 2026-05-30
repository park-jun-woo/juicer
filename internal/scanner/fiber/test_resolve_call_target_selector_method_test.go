//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_SelectorMethod 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

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
