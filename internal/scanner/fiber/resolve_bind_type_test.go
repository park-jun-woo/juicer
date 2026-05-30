//ff:func feature=scan type=test control=sequence
//ff:what resolveBindType — 바인딩 타입 해석 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveBindType_NoArgsOrNilInfo(t *testing.T) {
	if tn, f := resolveBindType(&ast.CallExpr{}, newEmptyInfo()); tn != "" || f != nil {
		t.Fatalf("no args: got %q %v", tn, f)
	}
	call := parseCall(t, "c.BodyParser(&req)")
	if tn, f := resolveBindType(call, nil); tn != "" || f != nil {
		t.Fatalf("nil info: got %q %v", tn, f)
	}
}

func TestResolveBindType_Resolved(t *testing.T) {
	src := `package m
type Req struct {
	Name string ` + "`json:\"name\"`" + `
}
func h() {
	var req Req
	parse(&req)
}
func parse(v interface{}) {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("m", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "parse" {
				call = c
			}
		}
		return true
	})
	if call == nil {
		t.Fatal("parse call not found")
	}
	tn, fields := resolveBindType(call, info)
	if tn != "Req" {
		t.Fatalf("typeName = %q, want Req", tn)
	}
	if len(fields) != 1 || fields[0].JSON != "name" {
		t.Fatalf("fields = %+v", fields)
	}
}
