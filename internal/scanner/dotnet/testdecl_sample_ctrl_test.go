//ff:type feature=scan type=test topic=dotnet
//ff:what sampleCtrl 테스트 보조 선언
package dotnet

const sampleCtrl = `
[ApiController]
[Route("api/[controller]")]
public class UsersController : ControllerBase {
    [HttpGet("{id}")]
    public ActionResult<UserDto> Get(int id) { return Ok(); }

    [HttpPost]
    public ActionResult<UserDto> Create([FromBody] CreateUserDto dto) { return Ok(); }
}
`
