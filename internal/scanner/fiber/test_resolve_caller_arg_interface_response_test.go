//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_InterfaceResponse 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArg_InterfaceResponse(t *testing.T) {

	src := `package m
type Resp struct {
	OK bool ` + "`json:\"ok\"`" + `
}
func h() {
	var r Resp
	sink(r)
}
func sink(v interface{}) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "sink" {
				arg = c.Args[0]
			}
		}
		return true
	})
	emptyIface := types.NewInterfaceType(nil, nil)
	res := resolveCallerArg(emptyIface, arg, info)
	if res.typeName != "Resp" {
		t.Fatalf("expected typeName Resp, got %+v", res)
	}
}
