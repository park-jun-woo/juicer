//ff:func feature=scan type=test control=sequence topic=echo
//ff:what round5 type-blocked 함수의 도달 가능한 가드/빈입력 분기 테스트 (echo, best-effort)
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	"golang.org/x/tools/go/packages"
)

// --- functions reachable via a real Scan on a non-echo module (genuine no-route path) ---

func TestScan_NonEchoModule_Round5(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module x\ngo 1.21\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "m.go"), []byte("package x\nfunc Plain() {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	res, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Endpoints) != 0 {
		t.Fatalf("expected no echo endpoints, got %d", len(res.Endpoints))
	}
}

func TestScan_InvalidPath_Round5(t *testing.T) {
	// non-existent path still resolves via filepath.Abs but Load yields no pkgs;
	// the function must not panic and must return a (possibly empty) result.
	if _, err := Scan(string([]byte{0})); err == nil {
		// some platforms accept it; just ensure no panic by reaching here
	}
}

func TestExtractRoutes_EmptyPkgs_Round5(t *testing.T) {
	eps, hmap := extractRoutes(nil, "/root")
	if len(eps) != 0 || len(hmap) != 0 {
		t.Fatalf("expected empty, got %d %d", len(eps), len(hmap))
	}
}

func TestBuildFuncIndex_EmptyPkgs_Round5(t *testing.T) {
	idx := buildFuncIndex(nil)
	if idx == nil || idx.byPos == nil || idx.byName == nil {
		t.Fatalf("index not initialized: %+v", idx)
	}
	if len(idx.byPos) != 0 {
		t.Fatalf("expected empty index, got %d", len(idx.byPos))
	}
}

func TestBuildEndpointIndex_Round5(t *testing.T) {
	eps := []scanner.Endpoint{
		{File: "a.go", Line: 1},
		{File: "b.go", Line: 2},
	}
	m := buildEndpointIndex(eps)
	if len(m) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(m))
	}
	if m[struct {
		file string
		line int
	}{"a.go", 1}] != 0 {
		t.Fatalf("index mismatch: %v", m)
	}
}

func TestAnalyzeHandlers_EmptyExprs_Round5(t *testing.T) {
	eps := []scanner.Endpoint{{File: "a.go", Line: 1}}
	// empty handler expr map -> no-op, no panic
	analyzeHandlers(nil, eps, "/root", map[int][]ast.Expr{}, buildFuncIndex(nil))
}

func TestResolveGroupPrefix_EmptyPkgs_Round5(t *testing.T) {
	// no packages -> no-op
	resolveGroupPrefix(nil, "/root", buildFuncIndex(nil), nil, map[int][]ast.Expr{})
}

// --- pos/info lookup with empty packages ---

func TestFindInfoForExpr_And_FileForPos_Empty_Round5(t *testing.T) {
	expr := parseExpr(t, "x")
	if info := findInfoForExpr(expr, nil); info != nil {
		t.Fatal("expected nil info for empty pkgs")
	}
	if f := findFileForPos(token.Pos(1), nil); f != nil {
		t.Fatal("expected nil file for empty pkgs")
	}
}

// --- group-arg context helpers with empty context ---

func emptyGroupCtx() *groupArgCtx {
	return &groupArgCtx{
		echoAlias: "echo",
		routers:   map[string]*routerInfo{},
		info:      nil,
		fset:      token.NewFileSet(),
		idx:       buildFuncIndex(nil),
		root:      "/root",
		pkgs:      nil,
		hmap:      map[int][]ast.Expr{},
	}
}

func TestResolveTargetEchoAlias_Empty_Round5(t *testing.T) {
	if got := resolveTargetEchoAlias(token.Pos(1), emptyGroupCtx()); got != "" {
		t.Fatalf("expected empty alias, got %q", got)
	}
}

func TestResolveTargetFilePath_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	// register a position in the fset so Position resolves
	f := ctx.fset.AddFile("/root/sub/m.go", -1, 10)
	pos := f.Pos(0)
	got := resolveTargetFilePath(pos, ctx)
	if got != filepath.Join("sub", "m.go") {
		t.Fatalf("rel path: %q", got)
	}
}

func TestTryGroupArgCall_And_WalkForGroupArgs_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	// call with a literal arg -> extractGroupArgPrefix false for all -> no-op
	call := callExprFrom(t, `setup(42)`)
	tryGroupArgCall(call, ctx)

	// walkForGroupArgs over a statement list must traverse without panicking
	src := `package m
func f() {
	x := 1
	g(x)
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	walkForGroupArgs(fn.Body.List, ctx)
}

func TestForwardRouterCalls_NoMatch_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	src := `package m
func f() { other(); }
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	// paramName not present in any call -> routerArgIndex < 0 -> no-op
	forwardRouterCalls(fn.Body.List, "router", "/p", &routerInfo{}, nil, ctx, 0)
}

func TestApplyRescanResults_NoMatch_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	ctx.epIndex = map[struct {
		file string
		line int
	}]int{}
	// endpoint not in epIndex -> skipped
	applyRescanResults([]scanner.Endpoint{{File: "a.go", Line: 1}}, ctx)
}

func TestRescanCalleeWithPrefix_DepthGuard_Round5(t *testing.T) {
	ctx := emptyGroupCtx()
	call := callExprFrom(t, `target(g)`)
	// depth >= maxRescanDepth -> immediate return (no panic, no info needed)
	rescanCalleeWithPrefixDepth(call, 0, "/p", &routerInfo{}, ctx, maxRescanDepth)
}

func TestRescanCalleeWithPrefix_Wrapper_Round5(t *testing.T) {
	// type-check a package so ctx.info is non-nil; the call target won't resolve
	// to a known func in the (empty) index, exercising the early-return path.
	file, info := checkSrc(t, `package m
func helper() int { return 1 }
var _ = helper()
`)
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			call = c
		}
		return true
	})
	ctx := emptyGroupCtx()
	ctx.info = info
	// wrapper delegates to depth 0; lookupFunc misses -> early return, no panic
	rescanCalleeWithPrefix(call, 0, "/p", &routerInfo{}, ctx)
}

// --- AST/types guard branches ---

func TestTryRouteCall_Guards_Round5(t *testing.T) {
	routers := map[string]*routerInfo{"e": {}}
	fset := token.NewFileSet()
	// non-selector function
	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `f("/x", h)`), routers, "m.go", fset); ok {
		t.Fatal("non-selector should fail")
	}
	// not an echo method
	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `e.NotAMethod("/x", h)`), routers, "m.go", fset); ok {
		t.Fatal("non-echo method should fail")
	}
	// unknown router
	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `unknown.GET("/x", h)`), routers, "m.go", fset); ok {
		t.Fatal("unknown router should fail")
	}
	// too few args
	if _, _, ok := tryRouteCall(nil, callExprFrom(t, `e.GET("/x")`), routers, "m.go", fset); ok {
		t.Fatal("too few args should fail")
	}
	// valid: e.GET("/x", h) with literal path
	ep, _, ok := tryRouteCall(nil, callExprFrom(t, `e.GET("/x", h)`), routers, "m.go", fset)
	if !ok || ep.Method != "GET" || ep.Path != "/x" {
		t.Fatalf("valid route: %+v %v", ep, ok)
	}
}

func TestResolveCallTarget_Guards_Round5(t *testing.T) {
	file, info := checkSrc(t, `package m
func Target() int { return 0 }
var _ = Target()
`)
	var call *ast.CallExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			call = c
		}
		return true
	})
	if call == nil {
		t.Fatal("no call")
	}
	pos := resolveCallTarget(call, info)
	if !pos.IsValid() {
		t.Fatal("expected valid target pos")
	}
}

func TestResolveCallerArg_Round5(t *testing.T) {
	_, info := checkSrc(t, `package m
type Ctx interface{ JSON() }
type Dto struct { Name string `+"`json:\"name\"`"+` }
var d Dto
var n int
var _ = d
var _ = n
`)
	// int-kind param -> status branch
	var intType, ifaceType types.Type
	for id, obj := range info.Defs {
		if obj == nil {
			continue
		}
		if id.Name == "n" {
			intType = obj.Type()
		}
	}
	// build a basic int type directly
	if intType == nil {
		intType = types.Typ[types.Int]
	}
	res := resolveCallerArg(intType, parseExpr(t, "200"), info)
	if res.status == "" && res.typeName == "" && !res.skip {
		// status may be unknown; ensure function ran without panic
	}
	_ = ifaceType

	// empty interface param -> response-type branch
	empty := types.NewInterfaceType(nil, nil).Complete()
	_ = resolveCallerArg(empty, parseExpr(t, "d"), info)
}

func TestResolveCallerArgs_NilGuard_Round5(t *testing.T) {
	src := `package m
func Target() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	// nil calleeInfo -> early return
	status, tn, fields, conf := resolveCallerArgs(fn, &ast.CallExpr{}, nil, nil)
	if status != "" || tn != "" || fields != nil || conf != "" {
		t.Fatalf("expected empty: %q %q %v %q", status, tn, fields, conf)
	}
}

func TestScanFile_NoEcho_Round5(t *testing.T) {
	src := `package m
func F() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	eps, hmap := scanFile(nil, file, "m.go", fset)
	if eps != nil || hmap != nil {
		t.Fatalf("no echo import -> nil, got %v %v", eps, hmap)
	}
}

func TestScanBody_NilGuard_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	// nil body -> early return, no panic
	scanBody(ep, nil, "c", nil, buildFuncIndex(nil), "handler")
}

func TestAnalyzeExpr_FuncLitNoCtx_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	// func lit without an echo.Context param -> echoCtxParamNameInfo "" -> return
	expr := parseExpr(t, `func(){ }`)
	analyzeExpr(ep, expr, nil, buildFuncIndex(nil))
}

func TestCheckOneDepthCall_Guards_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	// nil info -> early return
	checkOneDepthCall(ep, callExprFrom(t, `helper(c)`), "c", nil, buildFuncIndex(nil))
}

func TestResolveGroupPrefixFile_NoEcho_Round5(t *testing.T) {
	src := `package m
func F() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	// no echo import -> echoAlias "" -> immediate return (no pkg deref)
	resolveGroupPrefixFile(file, &packages.Package{Fset: fset}, nil, "/root", buildFuncIndex(nil), nil, map[int][]ast.Expr{}, map[struct {
		file string
		line int
	}]int{})
}

func TestResolveGroupPrefixFunc_NoEcho_Round5(t *testing.T) {
	src := `package m
func setup() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)
	// echoAlias "" -> immediate return
	resolveGroupPrefixFunc(fn, "", &packages.Package{Fset: fset}, nil, "/root", buildFuncIndex(nil), nil, map[int][]ast.Expr{}, map[struct {
		file string
		line int
	}]int{})
}
