//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestIsMapStringAny 테스트
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

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
