//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_TypesFallback 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestInferValueType_TypesFallback(t *testing.T) {

	src := `package m
var V = len("abc")
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	// find the len(...) call expr
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok && call == nil {
			call = c
			return false
		}
		return true
	})
	got := inferValueType(call, info)
	if got != "int" {
		t.Fatalf("types fallback: got %q, want int", got)
	}
}
