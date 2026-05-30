//ff:type feature=scan type=test topic=dotnet
//ff:what round5MinimalSource 테스트 보조 선언
package dotnet

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
