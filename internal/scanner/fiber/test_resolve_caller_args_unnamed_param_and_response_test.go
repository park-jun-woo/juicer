//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestResolveCallerArgs_UnnamedParamAndResponse 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_UnnamedParamAndResponse(t *testing.T) {

	src := `package m
type Out struct {
	V int ` + "`json:\"v\"`" + `
}
func write(interface{}, status int) {}
func h() {
	var o Out
	write(o, 200)
}
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
	var fnDecl *ast.FuncDecl
	var call *ast.CallExpr
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "write" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "write" {
				call = c
			}
		}
		return true
	})
	status, tn, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "200" {
		t.Errorf("status = %q, want 200", status)
	}
	if tn != "Out" {
		t.Errorf("typeName = %q, want Out", tn)
	}
}
