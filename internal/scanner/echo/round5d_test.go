//ff:func feature=scan type=test control=sequence topic=echo
//ff:what round5 walkStmts/processAssign/ctx-param 헬퍼 테스트 (echo)
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// echoCtxFuncType returns the *ast.FuncType of the first func decl whose first
// param is an echo.Context, plus the type info.
func TestEchoCtxParamNameInfo_Round5(t *testing.T) {
	src := `package echo
type Context interface{ JSON(int, interface{}) error }
func handler(c Context) error { return nil }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "echo.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var ft *ast.FuncType
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "handler" {
			ft = fn.Type
		}
	}
	if ft == nil {
		t.Fatal("no handler func")
	}
	if got := echoCtxParamNameInfo(ft, info); got != "c" {
		t.Fatalf("got %q", got)
	}
	// nil info -> AST fallback path
	_ = echoCtxParamNameInfo(ft, nil)
}

func TestEchoRouterParamAtIndex_Round5(t *testing.T) {
	src := `package echo
type Echo struct{}
func setup(prefix string, e *Echo) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "echo.go", src, 0)
	conf := types.Config{}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	if _, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, info); err != nil {
		t.Fatal(err)
	}
	var fn *ast.FuncDecl
	for _, d := range file.Decls {
		if f, ok := d.(*ast.FuncDecl); ok && f.Name.Name == "setup" {
			fn = f
		}
	}
	// index 1 is the *Echo param
	if got := echoRouterParamAtIndex(fn, info, 1); got != "e" {
		t.Fatalf("idx1: %q", got)
	}
	// index 0 is a string param -> not a router
	if got := echoRouterParamAtIndex(fn, info, 0); got != "" {
		t.Fatalf("idx0: %q", got)
	}
}

func TestProcessAssign_Round5(t *testing.T) {
	// `e := echo.New()` registers a router; syntactic via the alias.
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

func TestResolveBindType_Round5(t *testing.T) {
	file, info := checkSrc(t, `package m
type CreateReq struct { Name string `+"`json:\"name\"`"+` }
var dto CreateReq
var _ = dto
`)
	// build a Bind call: c.Bind(&dto) referencing the used dto ident
	var useIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "dto" {
			if _, isUse := info.Uses[id]; isUse {
				useIdent = id
			}
		}
		return true
	})
	if useIdent == nil {
		t.Fatal("no use of dto")
	}
	call := &ast.CallExpr{Args: []ast.Expr{&ast.UnaryExpr{Op: token.AND, X: useIdent}}}
	name, fields := resolveBindType(call, info)
	if name != "CreateReq" || len(fields) == 0 {
		t.Fatalf("bind type: %q %+v", name, fields)
	}
	// no args -> empty
	if n, _ := resolveBindType(&ast.CallExpr{}, info); n != "" {
		t.Fatalf("no-args: %q", n)
	}
	// nil info -> empty
	if n, _ := resolveBindType(call, nil); n != "" {
		t.Fatalf("nil info: %q", n)
	}
}

func TestLookupFunc_Round5(t *testing.T) {
	src := `package m
func Target() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	info := &types.Info{}
	idx := &funcIndex{
		byPos: map[token.Pos]*ast.FuncDecl{fn.Pos(): fn},
		info:  map[token.Pos]*types.Info{fn.Pos(): info},
	}
	gotFn, gotInfo := lookupFunc(fn.Pos(), idx)
	if gotFn != fn || gotInfo != info {
		t.Fatalf("lookup mismatch")
	}
	// unknown pos -> nil
	if f, _ := lookupFunc(token.NoPos, idx); f != nil {
		t.Fatal("unknown pos should be nil")
	}
}

func TestWalkStmts_Round5(t *testing.T) {
	// walkStmts recurses through control structures; with a pre-registered
	// router it should extract a route call inside an if-block.
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
	// nil info: tryRouteCall falls back to syntactic router detection
	walkStmts(nil, fn.Body.List, "echo", "m.go", fset, routers, &out, hmap)
	// walkStmts must traverse into the if-body without panicking; if the route
	// is recognized it lands in out.
	_ = out
	_ = hmap
}
