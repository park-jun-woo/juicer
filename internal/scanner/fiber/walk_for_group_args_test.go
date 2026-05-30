//ff:func feature=scan type=test control=iteration dimension=1
//ff:what walkForGroupArgs — 문 순회 라우터 구축 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"golang.org/x/tools/go/packages"
)

func TestWalkForGroupArgs_AllStmtTypes(t *testing.T) {
	src := `package m
func Setup() {
	app := fiber.New()
	app.Use(mw)
	{
		api := app.Group("/api")
		_ = api
	}
	if x := 1; x > 0 {
		v1 := app.Group("/v1")
		_ = v1
	} else {
		_ = app.Group("/else")
	}
	for i := 0; i < 1; i++ {
		_ = app.Group("/for")
	}
	for range items {
		_ = app.Group("/range")
	}
	switch x {
	case 1:
		_ = app.Group("/case")
	}
	switch t := y.(type) {
	default:
		_ = t
	}
	select {
	case <-done:
	default:
		doThing()
	}
	<-ch
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "m.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var fn *ast.FuncDecl
	for _, d := range file.Decls {
		if f, ok := d.(*ast.FuncDecl); ok {
			fn = f
		}
	}
	ctx := &groupArgCtx{
		fiberAlias: "fiber",
		routers:    map[string]*routerInfo{},
		idx:        &funcIndex{byPos: map[token.Pos]*ast.FuncDecl{}},
		info:       newEmptyInfo(),
		pkgs:       []*packages.Package{},
	}
	// exercises all statement-type branches without panic
	walkForGroupArgs(fn.Body.List, ctx)

	// app was registered via fiber.New(); api/v1 groups registered too
	if _, ok := ctx.routers["app"]; !ok {
		t.Fatalf("app not registered: %v", ctx.routers)
	}
	if _, ok := ctx.routers["api"]; !ok {
		t.Errorf("api group not registered from nested block")
	}
}
