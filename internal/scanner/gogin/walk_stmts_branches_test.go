//ff:func feature=scan type=test control=iteration dimension=1
//ff:what walkStmts — 모든 문 유형 라우트 등록 테스트
package gogin

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestWalkStmts_AllStmtTypesRoutes(t *testing.T) {
	src := `package m
func Setup() {
	r := gin.New()
	r.Use(mw)
	r.GET("/root", h)
	{
		r.POST("/block", h)
	}
	if x := 1; x > 0 {
		r.PUT("/if", h)
	} else {
		r.DELETE("/else", h)
	}
	for i := 0; i < 1; i++ {
		r.GET("/for", h)
	}
	for range items {
		r.GET("/range", h)
	}
	switch x {
	case 1:
		r.GET("/case", h)
	}
	switch t := y.(type) {
	default:
		r.GET("/tswitch", h)
	}
	select {
	case <-done:
		r.GET("/comm", h)
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
	walkStmts(fn.Body.List, "gin", "m.go", fset, routers, &out, hmap)

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
