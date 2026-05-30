//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallerArg_GinContext 테스트
package gogin

import (
	"go/ast"
	"go/types"
	rcaTok "go/token"
	rcaParser "go/parser"
	rcaImp "go/importer"
	rcaTypes "go/types"
	"testing"
)

func TestResolveCallerArg_GinContext(t *testing.T) {
	// Non-gin pointer type should not be gin context
	ty := types.NewPointer(types.Typ[types.String])
	r := resolveCallerArg(ty, &ast.Ident{Name: "c"}, nil)
	if r.skip {
		t.Fatal("should not skip for non-gin pointer")
	}
}


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
