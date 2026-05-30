//ff:type feature=scan type=test topic=dotnet
//ff:what round5ControllerSource 테스트 보조 선언
package dotnet

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
