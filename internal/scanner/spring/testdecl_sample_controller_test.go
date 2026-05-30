//ff:type feature=scan type=test topic=spring
//ff:what sampleController 테스트 보조 선언
package spring

const sampleController = `
@RestController
@RequestMapping("/users")
public class UserController {
    @GetMapping("/{id}")
    public UserDto get(@PathVariable("id") Long id) { return null; }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    public UserDto create(@RequestBody CreateUserDto dto) { return null; }
}
`
