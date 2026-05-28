//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_QueryParams -- [FromQuery] 파라미터 추출 테스트
package dotnet

import "testing"

var queryParamControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/[controller]")]
public class ProductsController : ControllerBase
{
    [HttpGet]
    public IActionResult Search([FromQuery] string keyword, [FromQuery] int page = 1)
    {
        return Ok();
    }
}
`

func TestScan_QueryParams(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/ProductsController.cs", queryParamControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Path != "/api/products" {
		t.Errorf("path: want /api/products, got %s", ep.Path)
	}
	if ep.Request == nil {
		t.Fatalf("expected request")
	}
	if len(ep.Request.Query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(ep.Request.Query))
	}
	if ep.Request.Query[0].Name != "keyword" {
		t.Errorf("query[0] name: want keyword, got %s", ep.Request.Query[0].Name)
	}
	if ep.Request.Query[1].Name != "page" {
		t.Errorf("query[1] name: want page, got %s", ep.Request.Query[1].Name)
	}
	if ep.Request.Query[1].Default != "1" {
		t.Errorf("query[1] default: want 1, got %s", ep.Request.Query[1].Default)
	}
}
