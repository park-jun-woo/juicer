//ff:type feature=scan type=test topic=spring
//ff:what 타 패키지 부모 클래스 상속 테스트용 소스 템플릿
package spring

var crossPkgControllerSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;
import com.example.dto.AlbumRequest;

@RestController
public class AlbumController {

    @PostMapping("/albums")
    public void create(@RequestBody AlbumRequest body) {}
}
`

var crossPkgChildDtoSource = `
package com.example.dto;

import com.example.audit.UserDateAuditPayload;

public class AlbumRequest extends UserDateAuditPayload {
    private String title;
}
`

var crossPkgParentDtoSource = `
package com.example.audit;

public class UserDateAuditPayload {
    private Long userId;
    private String updatedAt;
}
`
