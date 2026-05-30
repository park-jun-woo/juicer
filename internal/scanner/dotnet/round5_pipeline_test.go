//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what round5 파이프라인 함수 직접 호출 테스트 (build/extract/match)
package dotnet

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractMapGroupsFromFile_Round5(t *testing.T) {
	fi := csFileInfo(t, `
var app = builder.Build();
var api = app.MapGroup("/api");
`)
	groups := map[string]string{}
	extractMapGroupsFromFile(fi, groups)
	if groups["api"] != "/api" {
		t.Fatalf("expected api=/api, got %v", groups)
	}
}

func TestMatchMapGroupDeclaration_Round5(t *testing.T) {
	fi := csFileInfo(t, `
var users = api.MapGroup("/users");
`)
	groups := map[string]string{"api": "/api"}
	stmt := findAllByType(fi.root, "local_declaration_statement")[0]
	matchMapGroupDeclaration(stmt, fi, groups)
	if groups["users"] != "/api/users" {
		t.Fatalf("expected nested prefix, got %v", groups)
	}
}

func TestExtractMinimalAPIsFromFile_Round5(t *testing.T) {
	fi := csFileInfo(t, `
var app = builder.Build();
app.MapGet("/health", () => Results.Ok());
`)
	eps := extractMinimalAPIsFromFile(fi, map[string]string{})
	if len(eps) != 1 || eps[0].Path != "/health" {
		t.Fatalf("got %+v", eps)
	}
}

func TestMatchMapMethod_Round5(t *testing.T) {
	fi := csFileInfo(t, `app.MapGet("/x", () => Results.Ok());`)
	inv := findAllByType(fi.root, "invocation_expression")[0]
	ep, ok := matchMapMethod(inv, fi, map[string]string{})
	if !ok || ep.Method != "GET" || ep.Path != "/x" {
		t.Fatalf("got %+v %v", ep, ok)
	}
	// non-map invocation
	fi2 := csFileInfo(t, `Console.WriteLine("x");`)
	inv2 := findAllByType(fi2.root, "invocation_expression")[0]
	if _, ok := matchMapMethod(inv2, fi2, map[string]string{}); ok {
		t.Fatal("non-map invocation should not match")
	}
}

func TestExtractLambdaParams_And_Request_Round5(t *testing.T) {
	fi := csFileInfo(t, `app.MapPost("/u/{id}", (int id, [FromBody] CreateDto dto) => Results.Ok());`)
	inv := findAllByType(fi.root, "invocation_expression")[0]
	args := findChildByType(inv, "argument_list")
	argNodes := childrenOfType(args, "argument")
	req := extractLambdaRequest(argNodes, fi.src, "/u/{id}")
	if req == nil {
		t.Fatal("expected request from lambda")
	}
	if req.Body == nil || req.Body.TypeName != "CreateDto" {
		t.Fatalf("body: %+v", req.Body)
	}
	if len(req.PathParams) != 1 || req.PathParams[0].Name != "id" {
		t.Fatalf("path params: %+v", req.PathParams)
	}

	// extractLambdaRequest with too few args -> nil
	if got := extractLambdaRequest(argNodes[:1], fi.src, "/u"); got != nil {
		t.Fatalf("expected nil for single arg, got %+v", got)
	}

	// direct extractLambdaParams call
	lambda := findChildByType(argNodes[1], "lambda_expression")
	if lambda == nil {
		t.Fatal("no lambda expression")
	}
	lreq := extractLambdaParams(lambda, fi.src, "/u/{id}")
	if lreq == nil || lreq.Body == nil || lreq.Body.TypeName != "CreateDto" {
		t.Fatalf("extractLambdaParams: %+v", lreq)
	}
}

func TestMatchReturnType_Round5(t *testing.T) {
	// generic_name ActionResult<UserDto> sets returnType
	root, src := parseCS(t, `class C { ActionResult<UserDto> M() { return Ok(); } }`)
	var ep endpointInfo
	gens := findAllByType(root, "generic_name")
	if len(gens) == 0 {
		t.Fatal("no generic_name")
	}
	if !matchReturnType(gens[0], src, &ep) {
		t.Fatal("expected match for generic_name")
	}
	if ep.returnType != "UserDto" {
		t.Fatalf("returnType: %q", ep.returnType)
	}

	// predefined void
	root2, src2 := parseCS(t, `class C { void M() {} }`)
	pre := findAllByType(root2, "predefined_type")
	var ep2 endpointInfo
	if !matchReturnType(pre[0], src2, &ep2) {
		t.Fatal("void should match")
	}
}

func TestBuildAllEndpoints_Round5(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	controllers := collectControllers([]*fileInfo{fi})
	if len(controllers) == 0 {
		t.Fatal("no controllers collected")
	}
	eps, _ := buildAllEndpoints(controllers, "/abs")
	if len(eps) != 2 {
		t.Fatalf("expected 2 endpoints, got %d: %+v", len(eps), eps)
	}
	var _ scanner.Endpoint = eps[0]

	// also exercise buildControllerEndpoints directly
	ceps, _ := buildControllerEndpoints(controllers[0], "/abs", 0)
	if len(ceps) != 2 {
		t.Fatalf("buildControllerEndpoints: expected 2, got %d", len(ceps))
	}
}

func TestBuildControllerInfo_Round5(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	cls := findAllByType(fi.root, "class_declaration")[0]
	ci := buildControllerInfo(cls, fi)
	if ci.className != "UsersController" {
		t.Fatalf("className: %q", ci.className)
	}
	if ci.prefix != "api/users" {
		t.Fatalf("prefix: %q", ci.prefix)
	}
	if len(ci.endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(ci.endpoints))
	}
}
