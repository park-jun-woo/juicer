//ff:func feature=scan type=test control=sequence
//ff:what resolveCallerArg — caller 인자 상태/타입 해석 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArg_IntStatus(t *testing.T) {
	// int param + literal status arg -> status set
	intType := types.Typ[types.Int]
	arg := &ast.BasicLit{Kind: token.INT, Value: "201"}
	res := resolveCallerArg(intType, arg, newEmptyInfo())
	if res.status != "201" {
		t.Fatalf("expected status 201, got %+v", res)
	}
}

func TestResolveCallerArg_IntUnknownStatus(t *testing.T) {
	intType := types.Typ[types.Int]
	// non-literal status arg -> "(unknown)" -> empty result
	res := resolveCallerArg(intType, &ast.Ident{Name: "statusVar"}, newEmptyInfo())
	if res.status != "" {
		t.Fatalf("expected empty status, got %+v", res)
	}
}

func TestResolveCallerArg_InterfaceResponse(t *testing.T) {
	// empty interface param + a typed struct arg -> response type resolved
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

func TestResolveCallerArg_OtherType(t *testing.T) {
	// string param -> none of the branches -> empty result
	res := resolveCallerArg(types.Typ[types.String], &ast.BasicLit{Kind: token.STRING, Value: `"x"`}, newEmptyInfo())
	if res.status != "" || res.typeName != "" {
		t.Fatalf("expected empty result for string param, got %+v", res)
	}
}

func TestResolveCallerArg_InterfaceNoMatch(t *testing.T) {
	// empty interface param but arg has no resolvable type (nil info) ->
	// empty result (not skip).
	emptyIface := types.NewInterfaceType(nil, nil)
	res := resolveCallerArg(emptyIface, &ast.Ident{Name: "x"}, newEmptyInfo())
	if res.skip || res.typeName != "" {
		t.Fatalf("expected empty non-skip result, got %+v", res)
	}
}
