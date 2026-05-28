//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_MapGroup -- MapGroup prefix 전파 테스트
package dotnet

import "testing"

var mapGroupSource = `
var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

var api = app.MapGroup("/api");
var users = api.MapGroup("/users");
users.MapGet("/", () => Results.Ok());
users.MapPost("/", ([FromBody] CreateUserRequest req) => Results.Created());
`

func TestScan_MapGroup(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Program.cs", mapGroupSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if ep0.Path != "/api/users" {
		t.Errorf("ep0 path: want /api/users, got %s", ep0.Path)
	}

	ep1 := result.Endpoints[1]
	if ep1.Path != "/api/users" {
		t.Errorf("ep1 path: want /api/users, got %s", ep1.Path)
	}
	if ep1.Method != "POST" {
		t.Errorf("ep1 method: want POST, got %s", ep1.Method)
	}
}
