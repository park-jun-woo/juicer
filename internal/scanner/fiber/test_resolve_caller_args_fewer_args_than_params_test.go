//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestResolveCallerArgs_FewerArgsThanParams 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_FewerArgsThanParams(t *testing.T) {
	src := `package m
func need(a int, b int) {}
func h() { need(200) }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, parser.AllErrors)
	if err != nil {

		_ = err
	}
	conf := types.Config{Importer: importer.Default(), Error: func(error) {}}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	conf.Check("m", fset, []*ast.File{file}, info)

	var fnDecl *ast.FuncDecl
	var call *ast.CallExpr
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "need" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "need" {
				call = c
			}
		}
		return true
	})

	status, _, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "200" {
		t.Fatalf("expected status 200 from single arg, got %q", status)
	}
}
