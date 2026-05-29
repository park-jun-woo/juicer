//ff:type feature=scan type=test topic=dotnet
//ff:what 암묵 body 바인딩 테스트용 소스 템플릿
package dotnet

var implicitBodyControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    [HttpPost]
    public ActionResult<UserDto> Create(UserCreateViewModel model)
    {
        return Ok();
    }

    [HttpPut("{id}")]
    public ActionResult<UserDto> Edit(int id, UserCreateViewModel model)
    {
        return Ok();
    }
}
`

var implicitBodyDtoSource = `
namespace MyApp.Models;

public class UserCreateViewModel
{
    public string Name { get; set; }
    public string Email { get; set; }
}
`
