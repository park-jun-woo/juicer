//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestScan_Authorize -- [Authorize] 역할 추출 테스트
package dotnet

import "testing"

var authorizeControllerSource = `
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Authorization;

namespace MyApp.Controllers;

[ApiController]
[Route("api/[controller]")]
[Authorize(Roles = "Admin")]
public class AdminController : ControllerBase
{
    [HttpGet]
    public IActionResult GetAll()
    {
        return Ok();
    }

    [HttpDelete("{id}")]
    [Authorize(Roles = "SuperAdmin")]
    public IActionResult Delete(int id)
    {
        return Ok();
    }
}
`

func TestScan_Authorize(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Controllers/AdminController.cs", authorizeControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 2 {
		t.Fatalf("expected 2 endpoints, got %d", len(result.Endpoints))
	}

	ep0 := result.Endpoints[0]
	if len(ep0.Roles) != 1 || ep0.Roles[0] != "Admin" {
		t.Errorf("ep0 roles: want [Admin], got %v", ep0.Roles)
	}

	ep1 := result.Endpoints[1]
	if len(ep1.Roles) != 1 || ep1.Roles[0] != "SuperAdmin" {
		t.Errorf("ep1 roles: want [SuperAdmin], got %v", ep1.Roles)
	}
}
