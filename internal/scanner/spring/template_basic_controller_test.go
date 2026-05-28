//ff:type feature=scan type=test topic=spring
//ff:what 기본 컨트롤러 테스트용 소스 템플릿
package spring

var basicControllerSource = `
package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;
import com.example.demo.dto.CreateUserRequest;
import com.example.demo.dto.UserDto;
import java.util.List;

@RestController
@RequestMapping("/api/users")
public class UserController {

    @GetMapping
    public List<UserDto> getAll() {
        return null;
    }

    @GetMapping("/{id}")
    public UserDto getById(@PathVariable Long id) {
        return null;
    }

    @PostMapping
    public UserDto create(@RequestBody CreateUserRequest body) {
        return null;
    }
}
`

var basicCreateUserDtoSource = `
package com.example.demo.dto;

public class CreateUserRequest {
    @NotBlank
    private String name;

    @NotNull
    private String email;

    private int age;
}
`

var basicUserDtoSource = `
package com.example.demo.dto;

public class UserDto {
    private Long id;
    private String name;
    private String email;
}
`
