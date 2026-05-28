//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestScan_AllHTTPMethods -- 모든 HTTP 메서드 추출 테스트
package dotnet

import "testing"

var allMethodsControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/items")]
public class ItemsController : ControllerBase
{
    [HttpGet]
    public IActionResult GetAll() => Ok();

    [HttpPost]
    public IActionResult Create() => Ok();

    [HttpPut("{id}")]
    public IActionResult Update(int id) => Ok();

    [HttpDelete("{id}")]
    public IActionResult Delete(int id) => Ok();

    [HttpPatch("{id}")]
    public IActionResult Patch(int id) => Ok();
}
`

func TestScan_AllHTTPMethods(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/ItemsController.cs", allMethodsControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 5 {
		t.Fatalf("expected 5 endpoints, got %d", len(result.Endpoints))
	}

	expected := []struct {
		method string
		path   string
	}{
		{"GET", "/api/items"},
		{"POST", "/api/items"},
		{"PUT", "/api/items/{id}"},
		{"DELETE", "/api/items/{id}"},
		{"PATCH", "/api/items/{id}"},
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
