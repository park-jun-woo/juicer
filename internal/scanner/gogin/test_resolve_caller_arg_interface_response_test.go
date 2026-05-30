//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_InterfaceResponse 테스트
package gogin

import (
	"go/ast"
	rcaImp "go/importer"
	rcaParser "go/parser"
	rcaTok "go/token"
	rcaTypes "go/types"
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
	fset := rcaTok.NewFileSet()
	file, _ := rcaParser.ParseFile(fset, "m.go", src, 0)
	conf := rcaTypes.Config{Importer: rcaImp.Default()}
	info := &rcaTypes.Info{
		Types: map[ast.Expr]rcaTypes.TypeAndValue{},
		Defs:  map[*ast.Ident]rcaTypes.Object{},
		Uses:  map[*ast.Ident]rcaTypes.Object{},
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
	res := resolveCallerArg(rcaTypes.NewInterfaceType(nil, nil), arg, info)
	if res.typeName != "Resp" {
		t.Fatalf("expected Resp, got %+v", res)
	}
}
