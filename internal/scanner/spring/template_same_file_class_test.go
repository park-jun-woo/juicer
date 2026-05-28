//ff:type feature=scan type=test topic=spring
//ff:what 같은 파일 내부 클래스 테스트용 소스 템플릿
package spring

var sameFileInnerClassSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;

@RestController
public class AuthController {

    @PostMapping("/login")
    public void login(@RequestBody LoginParam param) {}

    public static class LoginParam {
        private String email;
        private String password;
    }
}
`
