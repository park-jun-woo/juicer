//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_MinimalAPI -- Minimal API 스캔 테스트
package dotnet

import "testing"

var minimalAPISource = `
var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/users", () => Results.Ok());
app.MapPost("/users", ([FromBody] CreateUserRequest req) => Results.Created());
app.MapGet("/users/{id}", (int id) => Results.Ok());
`

func TestScan_MinimalAPI(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Program.cs", minimalAPISource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Method != "GET" {
		t.Errorf("ep0 method: want GET, got %s", ep0.Method)
	}
	if ep0.Path != "/users" {
		t.Errorf("ep0 path: want /users, got %s", ep0.Path)
	}

	ep1 := result.Endpoints[1]
	if ep1.Method != "POST" {
		t.Errorf("ep1 method: want POST, got %s", ep1.Method)
	}
	if ep1.Request == nil || ep1.Request.Body == nil {
		t.Fatalf("ep1 expected body")
	}
	if ep1.Request.Body.TypeName != "CreateUserRequest" {
		t.Errorf("ep1 body type: want CreateUserRequest, got %s", ep1.Request.Body.TypeName)
	}

	ep2 := result.Endpoints[2]
	if ep2.Path != "/users/{id}" {
		t.Errorf("ep2 path: want /users/{id}, got %s", ep2.Path)
	}
	if ep2.Request == nil || len(ep2.Request.PathParams) != 1 {
		t.Fatalf("ep2 expected 1 path param")
	}
	if ep2.Request.PathParams[0].Name != "id" {
		t.Errorf("ep2 param name: want id, got %s", ep2.Request.PathParams[0].Name)
	}
}
