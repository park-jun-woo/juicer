//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_IdentUse 테스트
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
