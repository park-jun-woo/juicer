//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestResolveCallerArgs_StatusFromInt 테스트
package fiber

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveCallerArgs_StatusFromInt(t *testing.T) {
	src := `package m
func respond(c interface{}, status int) {}
func h() {
	respond(nil, 201)
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
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "respond" {
			fnDecl = fn
		}
	}
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "respond" {
				call = c
			}
		}
		return true
	})
	if fnDecl == nil || call == nil {
		t.Fatal("respond decl/call not found")
	}
	status, _, _, _ := resolveCallerArgs(fnDecl, call, info, info)
	if status != "201" {
		t.Fatalf("expected status 201, got %q", status)
	}
}
