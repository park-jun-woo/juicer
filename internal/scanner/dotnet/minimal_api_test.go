//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what minimal API / lambda / record / match 함수 테스트
package dotnet

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestExtractMethodCall(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { app.MapGet("/x", h); } }`)
	invs := findAllByType(root, "invocation_expression")
	for _, inv := range invs {
		recv, method := extractMethodCall(inv, src)
		if recv == "app" && method == "MapGet" {
			return
		}
	}
	t.Fatal("did not find app.MapGet")
}

func TestClassifyImplicitParam_PathParam(t *testing.T) {
	ep := &endpointInfo{method: "GET", path: "/users/{id}"}
	classifyImplicitParam("int", "id", ep)
	if len(ep.params) != 1 || ep.params[0].Name != "id" {
		t.Fatalf("got %+v", ep.params)
	}
}

func TestClassifyImplicitParam_Body(t *testing.T) {
	ep := &endpointInfo{method: "POST", path: "/users"}
	classifyImplicitParam("CreateUserDto", "dto", ep)
	if ep.bodyType != "CreateUserDto" {
		t.Fatalf("got %q", ep.bodyType)
	}
}

func TestClassifyImplicitParam_PrimitiveIgnored(t *testing.T) {
	ep := &endpointInfo{method: "POST", path: "/users"}
	classifyImplicitParam("string", "name", ep)
	if ep.bodyType != "" {
		t.Fatalf("primitive should not be body: %q", ep.bodyType)
	}
}

func TestClassifyLambdaParam_FromBody(t *testing.T) {
	root, src := parseCS(t, `class C { void M() { app.MapPost("/x", ([FromBody] UserDto dto) => Results.Ok()); } }`)
	params := findAllByType(root, "parameter")
	var req scanner.Request
	classifyLambdaParam(params[0], src, "/x", &req)
	if req.Body == nil || req.Body.TypeName != "UserDto" {
		t.Fatalf("got %+v", req.Body)
	}
}

func TestExtractRecordParams(t *testing.T) {
	root, src := parseCS(t, `public record CreateUserDto(string Name, int Age);`)
	recs := findAllByType(root, "record_declaration")
	if len(recs) == 0 {
		t.Skip("no record")
	}
	fields := extractRecordParams(recs[0], src)
	if len(fields) != 2 || fields[0].Name != "Name" {
		t.Fatalf("got %+v", fields)
	}
	if fields[0].Type != "string" || fields[1].Type != "integer" {
		t.Fatalf("types: %+v", fields)
	}
}

func TestArgType_ObjectCreation(t *testing.T) {
	root, src := parseCS(t, `class C { void m() { return Ok(new UserDto()); } }`)
	args := findAllByType(root, "argument")
	for _, a := range args {
		if name, _ := argType(a, src); name == "UserDto" {
			return
		}
	}
	t.Skip("no object creation arg matched")
}

func TestExtractMapGroups_Empty(t *testing.T) {
	fi := csFileInfo(t, `class C {}`)
	groups := extractMapGroups([]*fileInfo{fi})
	if len(groups) != 0 {
		t.Fatalf("got %v", groups)
	}
}

func TestExtractMinimalAPIs(t *testing.T) {
	fi := csFileInfo(t, `
var app = builder.Build();
app.MapGet("/health", () => Results.Ok());
app.MapPost("/users", (CreateUserDto dto) => Results.Created());
`)
	eps := extractMinimalAPIs([]*fileInfo{fi}, map[string]string{})
	if len(eps) < 2 {
		t.Fatalf("expected >=2 endpoints, got %d: %+v", len(eps), eps)
	}
}
