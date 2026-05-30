//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestWalkStmts_Round5 테스트
package echo

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkStmts_Round5(t *testing.T) {

	src := `package m
func f(cond bool) {
	if cond {
		e.GET("/users", handler)
	}
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}
	var fn *ast.FuncDecl
	for _, d := range file.Decls {
		if f, ok := d.(*ast.FuncDecl); ok {
			fn = f
		}
	}
	routers := map[string]*routerInfo{"e": {}}
	var out []scanner.Endpoint
	hmap := map[int][]ast.Expr{}

	walkStmts(nil, fn.Body.List, "echo", "m.go", fset, routers, &out, hmap)

	_ = out
	_ = hmap
}
