//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what dotnet Scan 파이프라인 분기 커버리지 보강 테스트 (round5)
package dotnet

import "testing"

// Minimal API exercising: StatusCode invocation, Results.Json/Created,
// FromHeader/FromQuery/FromBody, path params, DI types, route constraints.
var round5MinimalSource = `
var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();

app.MapGet("/items/{id:int}", (int id, ILogger<Program> logger, [FromHeader] string token) =>
{
    return Results.StatusCode(503, new ItemResponse());
});

app.MapPost("/items", ([FromBody] CreateItemRequest req, [FromQuery] bool draft) =>
{
    return Results.Created("/items/1", new ItemResponse());
});

app.MapGet("/items", () => Results.Ok(new List<ItemResponse>()));
`

func TestScan_Round5Minimal(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Program.cs", round5MinimalSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	get := result.Endpoints[0]
	if get.Path != "/items/{id}" {
		t.Errorf("path: want /items/{id}, got %s", get.Path)
	}
	if get.Request == nil || len(get.Request.PathParams) != 1 {
		t.Fatalf("expected 1 path param, got %+v", get.Request)
	}
	if get.Request.PathParams[0].Name != "id" {
		t.Errorf("pathparam: want id, got %s", get.Request.PathParams[0].Name)
	}
	if len(get.Request.Headers) != 1 || get.Request.Headers[0].Name != "token" {
		t.Errorf("header: want token, got %+v", get.Request.Headers)
	}
	// ILogger should be filtered out as DI type (not a path param)
	for _, p := range get.Request.PathParams {
		if p.Name == "logger" {
			t.Error("logger should be filtered as DI type")
		}
	}
	post := result.Endpoints[1]
	if post.Request == nil || post.Request.Body == nil {
		t.Fatalf("POST expected body, got %+v", post.Request)
	}
	if post.Request.Body.TypeName != "CreateItemRequest" {
		t.Errorf("POST body type: want CreateItemRequest, got %s", post.Request.Body.TypeName)
	}
	if len(post.Request.Query) != 1 || post.Request.Query[0].Name != "draft" {
		t.Errorf("POST query: want draft, got %+v", post.Request.Query)
	}
	list := result.Endpoints[2]
	_ = list
}

// Controller exercising: records, data annotations, multiple attributes,
// array/nullable/generic property types, authorize roles.
var round5RecordSource = `
namespace MyApp.Models;

public record CreateItemRequest(
    [Required] string Name,
    int? Quantity,
    List<string> Tags,
    decimal[] Prices);

public class ItemResponse
{
    public string Id { get; set; }
    public DateTime? CreatedAt { get; set; }
    public List<TagDto> Tags { get; set; }
}

public class TagDto
{
    public string Label { get; set; }
}
`

var round5ControllerSource = `
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Authorization;

namespace MyApp.Controllers;

[ApiController]
[Route("api/v1/[controller]")]
[Authorize(Roles = "Admin,User")]
public class WidgetsController : ControllerBase
{
    [HttpPost]
    public ActionResult<ItemResponse> Create([FromBody] CreateItemRequest req)
    {
        return Ok(new ItemResponse());
    }

    [HttpGet("{id}")]
    public ActionResult<List<ItemResponse>> List(int id)
    {
        return Ok(new List<ItemResponse>());
    }

    [HttpDelete("{id}")]
    public IActionResult Remove(int id)
    {
        return StatusCode(503, new ItemResponse());
    }
}
`

func TestScan_Round5Controller(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "Models/Models.cs", round5RecordSource)
	writeFile(t, dir, "Controllers/WidgetsController.cs", round5ControllerSource)

	result, err := Scan(dir)
	if err != nil {
		t.Fatalf("Scan error: %v", err)
	}
	if len(result.Endpoints) != 3 {
		t.Fatalf("expected 3 endpoints, got %d", len(result.Endpoints))
	}

	create := result.Endpoints[0]
	if create.Path != "/api/v1/widgets" {
		t.Errorf("create path: want /api/v1/widgets, got %s", create.Path)
	}
	if len(create.Roles) == 0 {
		t.Errorf("expected authorize roles propagated, got %+v", create.Roles)
	}
	if create.Request == nil || create.Request.Body == nil {
		t.Fatalf("create body missing: %+v", create.Request)
	}
	// record params should be resolved as DTO fields
	if len(create.Request.Body.Fields) == 0 {
		t.Errorf("record fields not resolved: %+v", create.Request.Body)
	}

	list := result.Endpoints[1]
	if list.Path != "/api/v1/widgets/{id}" {
		t.Errorf("list path: want /api/v1/widgets/{id}, got %s", list.Path)
	}
	if len(list.Responses) == 0 || list.Responses[0].Body != "array" {
		t.Errorf("list should return array: %+v", list.Responses)
	}

	remove := result.Endpoints[2]
	if len(remove.Responses) == 0 {
		t.Fatalf("remove should have StatusCode response: %+v", remove.Responses)
	}
	found503 := false
	for _, r := range remove.Responses {
		if r.Status == "503" {
			found503 = true
			if r.TypeName != "ItemResponse" {
				t.Errorf("503 type: want ItemResponse, got %s", r.TypeName)
			}
		}
	}
	if !found503 {
		t.Errorf("expected 503 StatusCode response: %+v", remove.Responses)
	}
}
