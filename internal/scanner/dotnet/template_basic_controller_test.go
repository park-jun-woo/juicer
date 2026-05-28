//ff:type feature=scan type=test topic=dotnet
//ff:what 기본 컨트롤러 테스트용 소스 템플릿
package dotnet

var basicControllerSource = `
using Microsoft.AspNetCore.Mvc;

namespace MyApp.Controllers;

[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase
{
    [HttpGet]
    public ActionResult<List<UserDto>> GetAll()
    {
        return Ok();
    }

    [HttpGet("{id}")]
    public ActionResult<UserDto> GetById(int id)
    {
        return Ok();
    }

    [HttpPost]
    public ActionResult<UserDto> Create([FromBody] CreateUserRequest request)
    {
        return Ok();
    }
}
`

var basicCreateUserDtoSource = `
using System.ComponentModel.DataAnnotations;

namespace MyApp.Models;

public class CreateUserRequest
{
    [Required]
    [StringLength(100)]
    public string Name { get; set; }

    [EmailAddress]
    public string Email { get; set; }

    [Range(0, 150)]
    public int Age { get; set; }
}
`

var basicUserDtoSource = `
namespace MyApp.Models;

public class UserDto
{
    public long Id { get; set; }
    public string Name { get; set; }
    public string Email { get; set; }
}
`
