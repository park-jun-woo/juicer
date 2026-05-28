//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_HeaderParam -- [FromHeader] 파라미터 추출 테스트
package dotnet

import "testing"

var headerParamControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/[controller]")]
public class AuthController : ControllerBase
{
    [HttpGet]
    public IActionResult GetProfile([FromHeader] string authorization)
    {
        return Ok();
    }
}
`

func TestScan_HeaderParam(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/AuthController.cs", headerParamControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(result.Endpoints))
	}

	ep := result.Endpoints[0]
	if ep.Request == nil {
		t.Fatalf("expected request")
	}
	if len(ep.Request.Headers) != 1 {
		t.Fatalf("expected 1 header, got %d", len(ep.Request.Headers))
	}
	if ep.Request.Headers[0].Name != "authorization" {
		t.Errorf("header name: want authorization, got %s", ep.Request.Headers[0].Name)
	}
}
