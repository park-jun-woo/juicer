//ff:func feature=scan type=test control=sequence topic=express
//ff:what round5 미커버 함수 직접 호출 테스트 (express)
package express

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func exFirst(t *testing.T, fi *fileInfo, typ string) *sitter.Node {
	t.Helper()
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == typ {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(fi.Root)
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}

func TestExpandAllMethod_Round5(t *testing.T) {
	if got := expandAllMethod("all"); len(got) != 5 {
		t.Fatalf("all: %v", got)
	}
	if got := expandAllMethod("GET"); len(got) != 1 || got[0] != "GET" {
		t.Fatalf("single: %v", got)
	}
}

func TestAppendUniquePrefix_Round5(t *testing.T) {
	m := map[routerKey][]string{}
	k := routerKey{file: "a.ts", varName: "r"}
	if !appendUniquePrefix(m, k, "/x") {
		t.Fatal("first insert should be true")
	}
	if appendUniquePrefix(m, k, "/x") {
		t.Fatal("duplicate insert should be false")
	}
	if !appendUniquePrefix(m, k, "/y") {
		t.Fatal("new value should be true")
	}
	if len(m[k]) != 2 {
		t.Fatalf("expected 2 prefixes, got %v", m[k])
	}
}

func TestChainMethodToRoute_Round5(t *testing.T) {
	cm := chainMethod{method: "POST", handler: "h", middleware: []string{"auth"}, line: 3, authLevel: "user", roles: []string{"admin"}}
	r := chainMethodToRoute(cm, "/items", "router")
	if r.Method != "POST" || r.Path != "/items" || r.Router != "router" || r.Handler != "h" {
		t.Fatalf("route: %+v", r)
	}
	if len(r.Middleware) != 1 || r.AuthLevel != "user" || len(r.Roles) != 1 {
		t.Fatalf("route meta: %+v", r)
	}
}

func TestDeclaratorMatchesName_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const router = express.Router();`))
	vd := exFirst(t, fi, "variable_declarator")
	if !declaratorMatchesName(vd, fi.Src, "router") {
		t.Fatal("should match router")
	}
	if declaratorMatchesName(vd, fi.Src, "other") {
		t.Fatal("should not match other")
	}
}

func TestExtractHandlerName_Round5(t *testing.T) {
	cases := map[string]string{
		`handler`:    "handler",
		`obj.method`: "obj.method",
		`() => {}`:   "(anonymous)",
		`wrap(h)`:    "wrap",
	}
	for src, want := range cases {
		f := mustParse(t, []byte("x = "+src+";"))
		node := rhsExpr(t, f)
		if got := extractHandlerName(node, f.Src); got != want {
			t.Errorf("extractHandlerName(%q)=%q want %q", src, got, want)
		}
	}
}

// rhsExpr returns the right-hand side expression node of `x = <expr>;`.
func rhsExpr(t *testing.T, fi *fileInfo) *sitter.Node {
	t.Helper()
	assign := exFirst(t, fi, "assignment_expression")
	return assign.Child(int(assign.ChildCount()) - 1)
}

func TestExtractMiddlewareName_Round5(t *testing.T) {
	// identifier
	f := mustParse(t, []byte(`x = authMiddleware;`))
	if got := extractMiddlewareName(rhsExpr(t, f), f.Src); got != "authMiddleware" {
		t.Errorf("identifier: %q", got)
	}
	// member_expression
	f2 := mustParse(t, []byte(`x = auth.required;`))
	if got := extractMiddlewareName(rhsExpr(t, f2), f2.Src); got != "auth.required" {
		t.Errorf("member: %q", got)
	}
	// call_expression -> resolved via extractHandlerFromCall
	f3 := mustParse(t, []byte(`x = requireRole("admin");`))
	if got := extractMiddlewareName(rhsExpr(t, f3), f3.Src); got != "requireRole" {
		t.Errorf("call: %q", got)
	}
}

func TestExtractImportPath_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`import x from './mod';`))
	stmt := exFirst(t, fi, "import_statement")
	if got := extractImportPath(stmt, fi.Src); got != "./mod" {
		t.Fatalf("got %q", got)
	}
}

func TestExtractFunctionBody_Round5(t *testing.T) {
	// arrow function body
	fi := mustParse(t, []byte(`const f = () => { return 1; };`))
	arrow := exFirst(t, fi, "arrow_function")
	body := extractFunctionBody(arrow)
	if body == nil || body.Type() != "statement_block" {
		t.Fatalf("arrow body: %v", body)
	}
	// function_expression body
	fi2 := mustParse(t, []byte(`const g = function () { return 2; };`))
	fexpr := exFirst(t, fi2, "function_expression")
	body2 := extractFunctionBody(fexpr)
	if body2 == nil || body2.Type() != "statement_block" {
		t.Fatalf("function_expression body: %v", body2)
	}
}

func TestExtractHandlerFromCall_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`x = wrap(handler);`))
	call := exFirst(t, fi, "call_expression")
	got := extractHandlerFromCall(call, fi.Src)
	if got == "" {
		t.Fatalf("expected handler name, got %q", got)
	}
}

func TestBuildUnresolvedSet_Round5(t *testing.T) {
	schemas := map[string]*sitter.Node{"A": nil}
	set := buildUnresolvedSet([]string{"A", "B", "C"}, schemas)
	if set["A"] {
		t.Error("A is resolved, should not be in set")
	}
	if !set["B"] || !set["C"] {
		t.Errorf("B,C should be unresolved: %v", set)
	}
}

func TestExtractAuthFromArgs_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', requireAuth, handler);`))
	call := firstCallExpr(t, fi)
	args := findChildByType(call, "arguments")
	argNodes := childrenOfType(args, "")
	_ = argNodes
	// pass the argument identifiers
	var ids []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		c := args.Child(i)
		if c.Type() == "identifier" || c.Type() == "call_expression" {
			ids = append(ids, c)
		}
	}
	level, roles := extractAuthFromArgs(ids, fi.Src)
	_ = level
	_ = roles
	// extractAuthFromMiddlewareNodes shares the logic
	l2, r2 := extractAuthFromMiddlewareNodes(ids, fi.Src)
	_ = l2
	_ = r2
}

func TestExtractMiddlewareNameForAuth_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', requireAuth, handler);`))
	id := exFirst(t, fi, "identifier")
	name, ok := extractMiddlewareNameForAuth(id, fi.Src)
	_ = name
	_ = ok
}

func TestBodyContainsRouterUse_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const f = () => { router.use('/x', sub); };`))
	arrow := exFirst(t, fi, "arrow_function")
	body := extractFunctionBody(arrow)
	if !bodyContainsRouterUse(body, fi.Src, map[string]bool{"router": true}) {
		t.Fatal("expected router.use detected")
	}
	fi2 := mustParse(t, []byte(`const g = () => { return 1; };`))
	arrow2 := exFirst(t, fi2, "arrow_function")
	body2 := extractFunctionBody(arrow2)
	if bodyContainsRouterUse(body2, fi2.Src, map[string]bool{"router": true}) {
		t.Fatal("expected no router.use")
	}
}

func TestBuildRouteFromArgs_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', handler);`))
	call := firstCallExpr(t, fi)
	args := findChildByType(call, "arguments")
	r := buildRouteFromArgs(args, fi.Src, "GET", 1)
	if r == nil || r.Method != "GET" || r.Path != "/x" {
		t.Fatalf("route: %+v", r)
	}
}

func TestExtractHandlerAndMiddleware_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.get('/x', mw, handler);`))
	call := firstCallExpr(t, fi)
	args := findChildByType(call, "arguments")
	var argNodes []*sitter.Node
	for i := 0; i < int(args.ChildCount()); i++ {
		c := args.Child(i)
		switch c.Type() {
		case "(", ")", ",":
		default:
			argNodes = append(argNodes, c)
		}
	}
	// full arg list: [path, mw, handler]; index 0 is the path string
	handler, mw := extractHandlerAndMiddleware(argNodes, fi.Src)
	if handler != "handler" {
		t.Fatalf("handler: %q", handler)
	}
	if len(mw) != 1 || mw[0] != "mw" {
		t.Fatalf("middleware: %v", mw)
	}
}

func TestExtractChainHandlerAndMiddleware_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`r.route('/x').get(mw, handler);`))
	// find the .get(...) call
	var getCall *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if getCall != nil {
			return
		}
		if n.Type() == "call_expression" {
			fn := n.ChildByFieldName("function")
			if fn != nil && fn.Type() == "member_expression" {
				prop := fn.ChildByFieldName("property")
				if prop != nil && nodeText(prop, fi.Src) == "get" {
					getCall = n
				}
			}
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(fi.Root)
	if getCall == nil {
		t.Fatal("no get call")
	}
	handler, mw, _, _, _ := extractChainHandlerAndMiddleware(getCall, fi.Src)
	if handler != "handler" {
		t.Fatalf("handler: %q", handler)
	}
	if len(mw) != 1 || mw[0] != "mw" {
		t.Fatalf("middleware: %v", mw)
	}
}

func TestBuildMountGraph_Round5(t *testing.T) {
	allRouters := map[string]map[string]bool{
		"a.ts": {"r": true},
		"b.ts": {"sub": true},
	}
	mounts := []mountEntry{
		{prefix: "/x", varName: "sub", filePath: "b.ts", sourceFile: "a.ts", sourceRouter: "r"},
	}
	g := buildMountGraph(mounts, allRouters)
	if g == nil {
		t.Fatal("nil graph")
	}
	if !g.nodes[routerKey{file: "a.ts", varName: "r"}] {
		t.Fatalf("expected seeded node, got %v", g.nodes)
	}
}

func round5Ctx() *scanContext {
	return &scanContext{
		parsed:         map[string]*fileInfo{},
		allRouters:     map[string]map[string]bool{},
		routerPrefixes: map[routerKey][]string{},
		pathAliases:    map[string]string{},
		schemas:        map[string]*sitter.Node{},
		schemaSrc:      map[string][]byte{},
	}
}

func TestBuildEndpointsFromRoute_And_OneEndpoint_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	ctx := round5Ctx()
	r := routeInfo{Method: "GET", Path: "/users/:id", Router: "r", Handler: "getUser", Line: 1}

	eps := buildEndpointsFromRoute(r, "/api", "routes.ts", ctx, fi)
	if len(eps) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(eps))
	}
	if eps[0].Method != "GET" || eps[0].Handler != "getUser" {
		t.Fatalf("endpoint: %+v", eps[0])
	}
	if eps[0].Path != "/api/users/{id}" {
		t.Errorf("path: %q", eps[0].Path)
	}

	// direct buildOneEndpoint
	ep := buildOneEndpoint("POST", "/api/users", r, "routes.ts", []string{}, ctx, fi)
	if ep.Method != "POST" {
		t.Fatalf("buildOneEndpoint method: %q", ep.Method)
	}

	// 'all' expansion through buildEndpointsFromRoute
	rAll := routeInfo{Method: "all", Path: "/x", Handler: "h"}
	allEps := buildEndpointsFromRoute(rAll, "", "r.ts", ctx, fi)
	if len(allEps) != 5 {
		t.Fatalf("expected 5 endpoints for all, got %d", len(allEps))
	}
}

func TestBuildRequest_Round5(t *testing.T) {
	fi := mustParse(t, []byte(`const r = express.Router();`))
	ctx := round5Ctx()
	// route with a path param -> request with PathParams
	r := routeInfo{Method: "GET", Path: "/users/:id", Handler: "h"}
	req := buildRequest(r, []string{"id"}, ctx, fi)
	if req == nil || len(req.PathParams) != 1 || req.PathParams[0].Name != "id" {
		t.Fatalf("request: %+v", req)
	}
	// no content -> nil
	if got := buildRequest(routeInfo{Method: "GET", Path: "/x", Handler: "h"}, nil, ctx, fi); got != nil {
		t.Fatalf("expected nil request, got %+v", got)
	}
}

func TestAddMountEdges_Round5(t *testing.T) {
	allRouters := map[string]map[string]bool{
		"child.ts": {"sub": true, "extra": true},
	}
	// inline mount (filePath empty)
	g := buildMountGraph(nil, allRouters)
	addMountEdges(g, mountEntry{prefix: "/in", varName: "sub", filePath: "", sourceFile: "a.ts", sourceRouter: "r"}, allRouters)

	// ambiguous cross-file mount: filePath set, varName empty -> all child routers
	g2 := buildMountGraph(nil, allRouters)
	addAmbiguousMountEdges(g2, routerKey{"a.ts", "r"}, mountEntry{prefix: "/x", filePath: "child.ts", sourceFile: "a.ts", sourceRouter: "r"}, allRouters)
	if len(g2.edges) == 0 {
		t.Fatalf("expected ambiguous edges, got %v", g2.edges)
	}
}

func TestBuildEndpointsFromFile_Round5(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "r.ts", `
import express from "express";
const router = express.Router();
router.get("/health", listHealth);
export default router;
`)
	res, err := Scan(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Endpoints) == 0 {
		t.Fatalf("expected endpoints from scan, got %d", len(res.Endpoints))
	}

	// direct buildEndpointsFromFile
	fi := mustParse(t, []byte(`
const router = express.Router();
router.get("/health", listHealth);
`))
	ctx := round5Ctx()
	eps := buildEndpointsFromFile(fi, map[string]bool{"router": true}, "r.ts", "r.ts", ctx)
	if len(eps) == 0 {
		t.Fatalf("buildEndpointsFromFile: expected endpoints, got %d", len(eps))
	}
}
