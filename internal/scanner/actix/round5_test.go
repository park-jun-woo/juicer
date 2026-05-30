//ff:func feature=scan type=test control=sequence topic=actix
//ff:what round5 미커버 함수 직접 호출 테스트 (actix)
package actix

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	sitter "github.com/smacker/go-tree-sitter"
)

func aParse(t *testing.T, src string) (*sitter.Node, []byte) {
	t.Helper()
	b := []byte(src)
	root, err := parseRust(b)
	if err != nil {
		t.Fatal(err)
	}
	return root, b
}

func aFi(t *testing.T, src string) *fileInfo {
	t.Helper()
	root, b := aParse(t, src)
	return &fileInfo{absPath: "/abs/m.rs", relPath: "m.rs", projectRoot: "/abs", src: b, root: root}
}

func aFirst(t *testing.T, root *sitter.Node, typ string) *sitter.Node {
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
	walk(root)
	if found == nil {
		t.Fatalf("no %s node", typ)
	}
	return found
}

// ---- pure helpers ----

func TestEnsureRequest_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	ensureRequest(ep)
	if ep.Request == nil {
		t.Fatal("request should be created")
	}
	prev := ep.Request
	ensureRequest(ep)
	if ep.Request != prev {
		t.Fatal("ensureRequest should not replace existing request")
	}
}

func TestApplyPathParams_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyPathParams(ep, "/users/{id}")
	if ep.Request == nil || len(ep.Request.PathParams) != 1 || ep.Request.PathParams[0].Name != "id" {
		t.Fatalf("path params: %+v", ep.Request)
	}
	// no params -> request stays nil
	ep2 := &scanner.Endpoint{}
	applyPathParams(ep2, "/static")
	if ep2.Request != nil {
		t.Fatalf("expected nil request, got %+v", ep2.Request)
	}
}

func TestApplyPrimitivePathType_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyPathParams(ep, "/users/{id}")
	applyPrimitivePathType(ep, "integer")
	if ep.Request.PathParams[0].Type != "integer" {
		t.Fatalf("type: %q", ep.Request.PathParams[0].Type)
	}
}

func TestApplyStructPathParams_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyStructPathParams(ep, []scanner.Field{{JSON: "id", Type: "integer"}, {JSON: "slug", Type: "string"}})
	if len(ep.Request.PathParams) != 2 {
		t.Fatalf("params: %+v", ep.Request.PathParams)
	}
}

func TestApplySerdeAttrs_Round5(t *testing.T) {
	name, nullable := applySerdeAttrs([]serdeAttr{{rename: "user_name"}, {hasDefault: true}}, "userName", false)
	if name != "user_name" || !nullable {
		t.Fatalf("got %q %v", name, nullable)
	}
	name2, nullable2 := applySerdeAttrs(nil, "x", false)
	if name2 != "x" || nullable2 {
		t.Fatalf("no attrs: %q %v", name2, nullable2)
	}
}

func TestDeduplicateBuilderRoutes_Round5(t *testing.T) {
	routes := []builderRoute{
		{method: "GET", path: "/a", handler: "h"},
		{method: "GET", path: "/a", handler: "h"},
		{method: "POST", path: "/a", handler: "h"},
	}
	got := deduplicateBuilderRoutes(routes)
	if len(got) != 2 {
		t.Fatalf("expected 2 unique, got %d", len(got))
	}
}

// ---- extractor application ----

func extractorTestSetup() (*scanner.Endpoint, structIndex, map[string][]scanner.Field) {
	return &scanner.Endpoint{}, structIndex{}, map[string][]scanner.Field{}
}

func TestApplyExtractors_AllKinds_Round5(t *testing.T) {
	for _, kind := range []string{"json", "query", "form", "path"} {
		ep, sIdx, cache := extractorTestSetup()
		ext := extractorInfo{kind: kind, typeName: "i64", rawType: "web::" + kind}
		applyExtractor(ep, ext, sIdx, cache)
	}
	// applyExtractors over a slice
	ep, sIdx, cache := extractorTestSetup()
	applyExtractors(ep, []extractorInfo{
		{kind: "json", typeName: "CreateReq"},
		{kind: "query", typeName: "Filter"},
	}, sIdx, cache)
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected json body: %+v", ep.Request)
	}
}

func TestApplyJSONExtractor_Round5(t *testing.T) {
	ep, sIdx, cache := extractorTestSetup()
	applyJSONExtractor(ep, extractorInfo{kind: "json", typeName: "CreateReq"}, sIdx, cache)
	if ep.Request.Body == nil || ep.Request.Body.TypeName != "CreateReq" {
		t.Fatalf("body: %+v", ep.Request.Body)
	}
}

func structIndexFor(t *testing.T, src string) structIndex {
	t.Helper()
	root, b := aParse(t, src)
	return buildStructIndex([]*fileInfo{{src: b, root: root}})
}

func TestApplyQueryExtractor_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	sIdx := structIndexFor(t, `struct Filter { page: i64, q: String }`)
	cache := map[string][]scanner.Field{}
	applyQueryExtractor(ep, extractorInfo{kind: "query", typeName: "Filter"}, sIdx, cache)
	if ep.Request == nil || len(ep.Request.Query) == 0 {
		t.Fatalf("expected query params: %+v", ep.Request)
	}
}

func TestApplyPathExtractor_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	applyPathParams(ep, "/u/{id}")
	sIdx := structIndex{}
	cache := map[string][]scanner.Field{}
	applyPathExtractor(ep, extractorInfo{kind: "path", typeName: "i64"}, sIdx, cache)
	if ep.Request == nil {
		t.Fatal("expected request")
	}
}

func TestApplyFormExtractor_Round5(t *testing.T) {
	ep := &scanner.Endpoint{}
	sIdx := structIndexFor(t, `struct FormData { title: String }`)
	cache := map[string][]scanner.Field{}
	applyFormExtractor(ep, extractorInfo{kind: "form", typeName: "FormData"}, sIdx, cache)
	if ep.Request == nil || len(ep.Request.FormFields) == 0 {
		t.Fatalf("expected form fields: %+v", ep.Request)
	}
}

// ---- AST: builder routes pipeline ----

const builderSrc = `
pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        web::resource("/users")
            .route(web::get().to(list_users))
            .route(web::post().to(create_user)),
    );
}
`

func TestExtractBuilderRoutes_Round5(t *testing.T) {
	fi := aFi(t, builderSrc)
	routes := extractBuilderRoutes(fi)
	if len(routes) == 0 {
		t.Fatalf("expected builder routes, got %d", len(routes))
	}
}

func TestBuildBuilderEndpoint_Round5(t *testing.T) {
	br := builderRoute{method: "GET", path: "/users/{id}", handler: "get_user"}
	ep := buildBuilderEndpoint(br, structIndex{}, map[string][]scanner.Field{}, map[string]*handlerInfo{})
	if ep.Method != "GET" || ep.Path != "/users/{id}" {
		t.Fatalf("endpoint: %+v", ep)
	}
	if ep.Request == nil || len(ep.Request.PathParams) != 1 {
		t.Fatalf("path params: %+v", ep.Request)
	}
}

func TestCollectRouteCalls_And_ToCalls_Round5(t *testing.T) {
	fi := aFi(t, builderSrc)
	var routes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectRouteCalls(n, fi.src, "/users", &routes)
	})
	if len(routes) == 0 {
		t.Fatalf("collectRouteCalls found nothing")
	}

	var toRoutes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectToCalls(n, fi.src, "/users", &toRoutes)
	})
}

func TestAppendRouteFromArgs_And_AppendToRoute_Round5(t *testing.T) {
	fi := aFi(t, `fn f() { web::get().to(handler); }`)
	args := aFirst(t, fi.root, "arguments")
	var routes []builderRoute
	appendRouteFromArgs(args, fi.src, "/x", &routes)
	var routes2 []builderRoute
	appendToRoute(args, fi.src, "/x", &routes2)
}

func TestExtractExtractors_And_Generic_Round5(t *testing.T) {
	fi := aFi(t, `async fn handler(body: web::Json<CreateReq>, id: web::Path<i64>) -> impl Responder { "" }`)
	fn := aFirst(t, fi.root, "function_item")
	exts := extractExtractors(fn, fi.src)
	if len(exts) == 0 {
		t.Fatalf("expected extractors, got %d", len(exts))
	}
	// buildGenericExtractor on a generic_type node
	gt := aFirst(t, fi.root, "generic_type")
	_ = buildGenericExtractor(gt, fi.src)
}

func TestApplyHandlerSignature_Round5(t *testing.T) {
	fi := aFi(t, `async fn handler(body: web::Json<CreateReq>) -> impl Responder { "" }`)
	fn := aFirst(t, fi.root, "function_item")
	ep := &scanner.Endpoint{}
	applyHandlerSignature(ep, fn, fi.src, structIndex{}, map[string][]scanner.Field{})
	if ep.Request == nil || ep.Request.Body == nil {
		t.Fatalf("expected body from handler signature: %+v", ep.Request)
	}
}

func TestCollectHandlerFuncs_Round5(t *testing.T) {
	fi := aFi(t, `async fn list_users() -> impl Responder { "" }`)
	index := map[string]*handlerInfo{}
	collectHandlerFuncs(fi, index)
	if _, ok := index["list_users"]; !ok {
		t.Fatalf("handler not collected: %v", index)
	}
}

// ---- macro routes pipeline ----

const macroSrc = `
#[get("/health")]
async fn health() -> impl Responder { "" }
`

func TestParseMacroAttribute_And_BuildMacroEndpoint_Round5(t *testing.T) {
	fi := aFi(t, macroSrc)
	attr := aFirst(t, fi.root, "attribute_item")
	mr := parseMacroAttribute(attr, fi.src)
	if mr == nil {
		t.Fatal("expected macro route")
	}
	if mr.method != "GET" || mr.path != "/health" {
		t.Fatalf("macro route: %+v", mr)
	}
	mr.file = fi
	mr.funcNode = aFirst(t, fi.root, "function_item")
	mr.handler = "health"
	ep := buildMacroEndpoint(*mr, fi, structIndex{}, map[string][]scanner.Field{})
	if ep.Method != "GET" || ep.Path != "/health" {
		t.Fatalf("endpoint: %+v", ep)
	}
}

func TestAppendMacroRoutes_And_ConsumeMacroChild_Round5(t *testing.T) {
	fi := aFi(t, macroSrc)
	fn := aFirst(t, fi.root, "function_item")
	var routes, pending []macroRoute
	// seed a pending route then attach to function
	pending = append(pending, macroRoute{method: "GET", path: "/health"})
	routes = appendMacroRoutes(routes, pending, fn, fi, "health")
	if len(routes) == 0 {
		t.Fatalf("expected macro routes attached")
	}

	// consumeMacroChild over the top-level children
	var r2, p2 []macroRoute
	for i := 0; i < int(fi.root.ChildCount()); i++ {
		r2, p2 = consumeMacroChild(fi.root.Child(i), fi, r2, p2)
	}
}

// ---- scope pipeline ----

func TestCaptureScope_And_ApplyScope_Round5(t *testing.T) {
	fi := aFi(t, `fn config(cfg: &mut web::ServiceConfig) { cfg.service(web::scope("/api").service(list_users)); }`)
	var scopes []scopeInfo
	walkNodes(fi.root, func(n *sitter.Node) {
		captureScope(n, fi.src, &scopes)
	})

	endpoints := []scanner.Endpoint{{Method: "GET", Path: "/users", Handler: "list_users"}}
	applyScopePrefixes([]*fileInfo{fi}, endpoints)
	applyScopeToEndpoints(scopeInfo{prefix: "/api", handlers: []string{"list_users"}}, endpoints)
	if endpoints[0].Path != "/api/users" {
		t.Fatalf("scope not applied: %q", endpoints[0].Path)
	}
}

func TestCollectServiceCallsAndHandlers_Round5(t *testing.T) {
	fi := aFi(t, `fn config(cfg: &mut web::ServiceConfig) { cfg.service(web::scope("/api").service(h1).service(h2)); }`)
	var routes []builderRoute
	walkNodes(fi.root, func(n *sitter.Node) {
		collectServiceCalls(n, fi.src, "/api", &routes)
		collectTopLevelServiceCall(n, fi, &routes)
	})

	scopeCall := aFirst(t, fi.root, "call_expression")
	_ = collectServiceHandlers(scopeCall, fi.src)

	// appendServiceCallHandlers
	_ = appendServiceCallHandlers(fi.root, fi.src, nil)
}

func TestCaptureHelpers_Round5(t *testing.T) {
	fi := aFi(t, `fn f() { web::scope("/api"); }`)
	call := aFirst(t, fi.root, "call_expression")
	var rootName string
	captureCallRoot(call, fi.src, &rootName)

	var arg string
	walkNodes(fi.root, func(n *sitter.Node) {
		captureScopedCallArg(n, fi.src, "scope", &arg)
	})
}

func TestConsumeFieldChild_Round5(t *testing.T) {
	fi := aFi(t, `struct S { #[serde(rename = "n")] name: String }`)
	var fields []scanner.Field
	var pending []serdeAttr
	field := aFirst(t, fi.root, "field_declaration")
	for i := 0; i < int(field.ChildCount()); i++ {
		fields, pending = consumeFieldChild(field.Child(i), fi.src, fields, pending)
	}
}

func TestCaptureResponse_Round5(t *testing.T) {
	fi := aFi(t, `async fn h() -> impl Responder { HttpResponse::Ok().json(user) }`)
	block := aFirst(t, fi.root, "block")
	ctx := &responseCtx{
		block:     block,
		src:       fi.src,
		sIdx:      structIndex{},
		cache:     map[string][]scanner.Field{},
		seen:      map[string]bool{},
		responses: nil,
	}
	walkNodes(block, func(n *sitter.Node) {
		captureResponse(n, ctx)
	})
}

func TestDetectResponseKind_Round5(t *testing.T) {
	fi := aFi(t, `async fn h() -> impl Responder { HttpResponse::Ok().json(user) }`)
	sid := aFirst(t, fi.root, "scoped_identifier")
	_ = detectResponseKind(sid, fi.src)
}

func TestExtractFirstStringArg_Round5(t *testing.T) {
	fi := aFi(t, `fn f() { web::resource("/users"); }`)
	call := aFirst(t, fi.root, "call_expression")
	if got := extractFirstStringArg(call, fi.src); got != "/users" {
		t.Fatalf("got %q", got)
	}
}
