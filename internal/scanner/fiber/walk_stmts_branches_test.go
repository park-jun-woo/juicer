//ff:func feature=scan type=test control=iteration dimension=1
//ff:what walkStmts — 모든 문 유형 순회 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWalkStmts_AllStmtTypes(t *testing.T) {
	src := `package m
func Setup() {
	app := fiber.New()
	app.Use(mw)
	app.Get("/root", h)
	{
		app.Post("/block", h)
	}
	if x := 1; x > 0 {
		app.Put("/if", h)
	} else {
		app.Delete("/else", h)
	}
	for i := 0; i < 1; i++ {
		app.Get("/for", h)
	}
	for range items {
		app.Get("/range", h)
	}
	switch x {
	case 1:
		app.Get("/case", h)
	}
	switch t := y.(type) {
	default:
		app.Get("/tswitch", h)
	}
	select {
	case <-done:
		app.Get("/comm", h)
	default:
	}
	<-ch
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "/proj/m.go", src, 0)
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
	var out []scanner.Endpoint
	hmap := map[int][]ast.Expr{}
	walkStmts(fn.Body.List, "fiber", "m.go", fset, routers, &out, hmap)

	paths := map[string]bool{}
	for _, ep := range out {
		paths[ep.Path] = true
	}
	for _, want := range []string{"/root", "/block", "/if", "/else", "/for", "/range", "/case", "/tswitch", "/comm"} {
		if !paths[want] {
			t.Errorf("missing route %s; got %v", want, paths)
		}
	}
}
