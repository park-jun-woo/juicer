//ff:func feature=scan type=test control=iteration dimension=1 topic=echo
//ff:what TestProcessAssign_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestProcessAssign_Round5(t *testing.T) {

	src := `package m
func f() {
	e := echo.New()
	g := e.Group("/api")
	_ = g
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
	routers := map[string]*routerInfo{}
	for _, stmt := range fn.Body.List {
		if as, ok := stmt.(*ast.AssignStmt); ok {
			processAssign(nil, as, "echo", routers)
		}
	}
	if _, ok := routers["e"]; !ok {
		t.Fatalf("e not registered: %v", routers)
	}
	if ri, ok := routers["g"]; !ok || ri.prefix != "/api" {
		t.Fatalf("g group prefix: %+v", routers["g"])
	}
}
