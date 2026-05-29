//ff:type feature=scan type=test topic=dotnet
//ff:what bare IActionResult 본문 분석 테스트용 소스 템플릿
package dotnet

var bodyResponseControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/[controller]")]
public class ItemsController : ControllerBase
{
    [HttpGet]
    public IActionResult GetAll()
    {
        return Ok(new ItemResponse());
    }

    [HttpPost]
    public IActionResult Create([FromBody] CreateItemRequest request)
    {
        if (!ModelState.IsValid)
        {
            return UnprocessableEntity(new ErrorResponse());
        }
        return Ok(new ResponseViewModel<ItemResponse>(item));
    }
}
`

var bodyResponseModelsSource = `
namespace MyApp.Models;

public class ItemResponse
{
    public long Id { get; set; }
    public string Title { get; set; }
}

public class ResponseViewModel<T>
{
    public bool Success { get; set; }
    public string Message { get; set; }
}

public class ErrorResponse
{
    public string Error { get; set; }
}

public class CreateItemRequest
{
    public string Title { get; set; }
}
`
