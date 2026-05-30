//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestResolveCallerArgs_NamedParamsDefs 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_NamedParamsDefs(t *testing.T) {

	src := `package m
func send(code int) {}
func h() { send(404) }
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
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "send" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "send" {
				call = c
			}
		}
		return true
	})
	status, _, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "404" {
		t.Fatalf("expected 404, got %q", status)
	}
}
