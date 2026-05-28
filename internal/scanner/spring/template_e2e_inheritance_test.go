//ff:type feature=scan type=test topic=spring
//ff:what E2E 상속 테스트용 소스 템플릿
package spring

var inheritanceControllerSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;
import com.example.dto.ChildDto;

@RestController
public class TestController {

    @PostMapping("/test")
    public void create(@RequestBody ChildDto body) {}
}
`

var inheritanceParentDtoSource = `
package com.example.dto;

public class BaseDto {
    private Long id;
    private String createdAt;
}
`

var inheritanceChildDtoSource = `
package com.example.dto;

public class ChildDto extends BaseDto {
    @NotBlank
    private String name;
}
`
