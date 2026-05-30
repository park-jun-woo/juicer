//ff:func feature=scan type=test control=sequence
//ff:what TestProcessAssign_FiberNew 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestProcessAssign_FiberNew(t *testing.T) {
	routers := map[string]*routerInfo{}
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "app"}},
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "fiber"},
					Sel: &ast.Ident{Name: "New"},
				},
			},
		},
		Tok: token.DEFINE,
	}

	processAssign(stmt, "fiber", routers)
	if _, ok := routers["app"]; !ok {
		t.Fatal("expected app in routers")
	}
}

func assignStmts(t *testing.T, src string) []*ast.AssignStmt {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parserParseFile(fset, src)
	if err != nil {
		t.Fatal(err)
	}
	var out []*ast.AssignStmt
	ast.Inspect(file, func(n ast.Node) bool {
		if a, ok := n.(*ast.AssignStmt); ok {
			out = append(out, a)
		}
		return true
	})
	return out
}

func TestProcessAssign_GroupAndSkips(t *testing.T) {
	src := `package m
func f() {
	app := fiber.New()
	api := app.Group("/api")
	x := 5            // rhs not a call -> skip
	y := compute()    // call but fun not selector -> skip
	_ = app.Group("/z") // lhs not ident -> skip
}
`
	routers := map[string]*routerInfo{}
	for _, a := range assignStmts(t, src) {
		processAssign(a, "fiber", routers)
	}
	if _, ok := routers["app"]; !ok {
		t.Fatal("app not registered")
	}
	if ri, ok := routers["api"]; !ok || ri.prefix != "/api" {
		t.Fatalf("api group not registered with prefix: %+v", routers["api"])
	}
}

func TestProcessAssign_GroupUnknownParent(t *testing.T) {
	src := `package m
func f() {
	sub := unknown.Group("/x")
}
`
	routers := map[string]*routerInfo{}
	for _, a := range assignStmts(t, src) {
		processAssign(a, "fiber", routers)
	}
	if _, ok := routers["sub"]; ok {
		t.Fatal("group with unknown parent should not register")
	}
}

func TestProcessAssign_MismatchedLhsRhs(t *testing.T) {
	// more Rhs than Lhs -> break
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.Ident{Name: "a"}},
		Rhs: []ast.Expr{
			&ast.BasicLit{Kind: token.INT, Value: "1"},
			&ast.CallExpr{Fun: &ast.SelectorExpr{X: &ast.Ident{Name: "fiber"}, Sel: &ast.Ident{Name: "New"}}},
		},
		Tok: token.DEFINE,
	}
	routers := map[string]*routerInfo{}
	processAssign(stmt, "fiber", routers)
	// second rhs (the fiber.New) is beyond Lhs len -> break, so not registered
	if len(routers) != 0 {
		t.Fatalf("expected no routers, got %v", routers)
	}
}

func parserParseFile(fset *token.FileSet, src string) (*ast.File, error) {
	return parser.ParseFile(fset, "m.go", src, 0)
}
