//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractMinimalAPIs 테스트
package dotnet

import "testing"

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
