//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestScan_ActionLevelRoute -- 액션 레벨 [Route] 합성 + version 토큰 정규화 테스트
package dotnet

import "testing"

var actionRouteControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/v{version:apiVersion}/users")]
public class UsersController : ControllerBase
{
    [HttpPost]
    [Route("login")]
    public IActionResult Login() => Ok();

    [HttpPost]
    [Route("logout")]
    public IActionResult Logout() => Ok();

    [HttpGet("{id}")]
    public IActionResult Info(int id) => Ok();
}
`

func TestScan_ActionLevelRoute(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/UsersController.cs", actionRouteControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	expected := []struct {
		method string
		path   string
	}{
		{"POST", "/api/v{version}/users/login"},
		{"POST", "/api/v{version}/users/logout"},
		{"GET", "/api/v{version}/users/{id}"},
	}
	for i, want := range expected {
		got := result.Endpoints[i]
		if got.Method != want.method {
			t.Errorf("ep[%d] method: want %s, got %s", i, want.method, got.Method)
		}
		if got.Path != want.path {
			t.Errorf("ep[%d] path: want %s, got %s", i, want.path, got.Path)
		}
	}
}
