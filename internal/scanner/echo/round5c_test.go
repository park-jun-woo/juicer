//ff:func feature=scan type=test control=sequence topic=echo
//ff:what round5 AST/types 핸들러·라우터 헬퍼 테스트 (echo)
package echo

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

// checkEchoPkg type-checks an in-memory package at the echo import path so that
// named types resolve with the echo package suffix.
func checkEchoPkg(t *testing.T, src string) (*ast.File, *types.Info, *types.Package) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "echo.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: map[ast.Expr]types.TypeAndValue{},
		Defs:  map[*ast.Ident]types.Object{},
		Uses:  map[*ast.Ident]types.Object{},
	}
	pkg, err := conf.Check("github.com/labstack/echo/v4", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatal(err)
	}
	return file, info, pkg
}

func TestRouterArgIndex_Round5(t *testing.T) {
	call := callExprFrom(t, `register(e, g, x)`)
	if got := routerArgIndex(call, "g"); got != 1 {
		t.Fatalf("got %d", got)
	}
	if got := routerArgIndex(call, "missing"); got != -1 {
		t.Fatalf("missing: %d", got)
	}
}

func TestParamFieldAtIndex_Round5(t *testing.T) {
	src := `package m
func F(a int, b, c string) {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, 0)
	fn := file.Decls[0].(*ast.FuncDecl)
	_, name0 := paramFieldAtIndex(fn.Type.Params, 0)
	if name0 != "a" {
		t.Fatalf("idx0: %q", name0)
	}
	_, name2 := paramFieldAtIndex(fn.Type.Params, 2)
	if name2 != "c" {
		t.Fatalf("idx2: %q", name2)
	}
	if f, _ := paramFieldAtIndex(fn.Type.Params, 9); f != nil {
		t.Fatal("out-of-range should be nil")
	}
}

func TestHandleForm_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callExprFrom(t, `c.FormValue("title")`)
	handleForm(ep, call)
	if ep.Request == nil || len(ep.Request.FormFields) != 1 || ep.Request.FormFields[0].Name != "title" {
		t.Fatalf("form: %+v", ep.Request)
	}
	handleForm(ep, call) // duplicate
	if len(ep.Request.FormFields) != 1 {
		t.Fatalf("dup: %+v", ep.Request.FormFields)
	}
}

func TestHandleFile_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callExprFrom(t, `c.FormFile("upload")`)
	handleFile(ep, call)
	if ep.Request == nil || len(ep.Request.Files) != 1 || ep.Request.Files[0].Name != "upload" {
		t.Fatalf("file: %+v", ep.Request)
	}
}

func TestHandleResponse_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
type UserDto struct { Name string `+"`json:\"name\"`"+` }
var u UserDto
var _ = u
`)
	call := callExprFrom(t, `c.JSON(200, u)`)
	ep := &scanner.Endpoint{}
	handleResponse(ep, call, "json", info, "handler")
	if len(ep.Responses) != 1 {
		t.Fatalf("responses: %+v", ep.Responses)
	}
	if ep.Responses[0].Kind != "json" {
		t.Fatalf("kind: %q", ep.Responses[0].Kind)
	}
}

func TestHandleBind_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	call := callExprFrom(t, `c.Bind(&dto)`)
	handleBind(ep, call, "Bind", nil)
	if ep.Request == nil || ep.Request.Body == nil || ep.Request.Body.VarName != "dto" {
		t.Fatalf("bind: %+v", ep.Request)
	}
	// second bind ignored
	handleBind(ep, callExprFrom(t, `c.Bind(&other)`), "Bind", nil)
	if ep.Request.Body.VarName != "dto" {
		t.Fatalf("second bind overwrote: %+v", ep.Request.Body)
	}
}

func TestBuildMapField_And_ExtractMapFields_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
var M = map[string]any{"name": "x", "age": 5}
`)
	var mapLit *ast.CompositeLit
	for e := range info.Types {
		if cl, ok := e.(*ast.CompositeLit); ok {
			if _, isMap := cl.Type.(*ast.MapType); isMap {
				mapLit = cl
			}
		}
	}
	if mapLit == nil {
		t.Fatal("no map literal")
	}
	fields := extractMapFields(mapLit, info)
	if len(fields) != 2 {
		t.Fatalf("fields: %+v", fields)
	}
	// buildMapField on a non-KV returns nil
	if buildMapField(&ast.BasicLit{}, info) != nil {
		t.Fatal("non-KV should be nil")
	}
}

func TestResolveResponseType_And_ExprType_Round5(t *testing.T) {
	file, info := checkSrc(t, `package m
type UserDto struct { Name string `+"`json:\"name\"`"+` }
var u UserDto
var _ = u
`)
	// find the use of u
	var useIdent *ast.Ident
	ast.Inspect(file, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok && id.Name == "u" {
			if _, isUse := info.Uses[id]; isUse {
				useIdent = id
			}
		}
		return true
	})
	if useIdent == nil {
		t.Fatal("no use of u")
	}
	name, fields := resolveExprType(useIdent, info)
	if name != "UserDto" || len(fields) == 0 {
		t.Fatalf("exprType: %q %+v", name, fields)
	}
	tn, flds, conf := resolveResponseType(useIdent, info)
	if tn != "UserDto" || len(flds) == 0 || conf != "full" {
		t.Fatalf("responseType: %q %+v %q", tn, flds, conf)
	}
	// nil info -> empty
	if n, _, _ := resolveResponseType(useIdent, nil); n != "" {
		t.Fatalf("nil info: %q", n)
	}
}

func TestIsEchoRouterTypeInfo_Round5(t *testing.T) {
	_, _, pkg := checkEchoPkg(t, `package echo
type Echo struct{}
type Group struct{}
type Other struct{}
`)
	echoT := pkg.Scope().Lookup("Echo").Type()
	if !isEchoRouterTypeInfo(echoT) {
		t.Fatal("Echo should be a router type")
	}
	// pointer to Group
	grpT := pkg.Scope().Lookup("Group").Type()
	if !isEchoRouterTypeInfo(types.NewPointer(grpT)) {
		t.Fatal("*Group should be a router type")
	}
	otherT := pkg.Scope().Lookup("Other").Type()
	if isEchoRouterTypeInfo(otherT) {
		t.Fatal("Other should not be a router type")
	}
}

func TestRegisterParams_Round5(t *testing.T) {
	src := `package m
import "github.com/labstack/echo/v4"
func setup(e *echo.Echo, g *echo.Group) {}
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
	routers := map[string]*routerInfo{}
	registerParams(fn, "echo", routers)
	if _, ok := routers["e"]; !ok {
		t.Fatalf("e not registered: %v", routers)
	}
	if _, ok := routers["g"]; !ok {
		t.Fatalf("g not registered: %v", routers)
	}
}

func TestExtractGroupArgPrefix_Round5(t *testing.T) {
	// variable argument resolved from routers map
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{"authGroup": {prefix: "/auth"}},
	}
	arg := exprFrom(t, `authGroup`)
	prefix, ri, ok := extractGroupArgPrefix(arg, ctx)
	if !ok || prefix != "/auth" || ri == nil {
		t.Fatalf("var group: %q %v", prefix, ok)
	}
	// non-group arg -> false
	if _, _, ok := extractGroupArgPrefix(exprFrom(t, `42`), ctx); ok {
		t.Fatal("literal should not be a group")
	}
}

// callExprFrom parses an expression and returns it as *ast.CallExpr.
func callExprFrom(t *testing.T, src string) *ast.CallExpr {
	t.Helper()
	expr := parseExpr(t, src)
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		t.Fatalf("not a call: %s", src)
	}
	return call
}

func exprFrom(t *testing.T, src string) ast.Expr {
	t.Helper()
	return parseExpr(t, src)
}
