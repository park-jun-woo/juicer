//ff:type feature=scan type=test topic=spring
//ff:what 상수 defaultValue 해석 테스트용 소스 템플릿
package spring

var constDefaultControllerSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;
import com.example.constants.AppConstants;
import java.util.List;

@RestController
@RequestMapping("/api/posts")
public class PostController {

    @GetMapping
    public List<String> list(
            @RequestParam(defaultValue = AppConstants.DEFAULT_PAGE_NUMBER) int page,
            @RequestParam(defaultValue = AppConstants.DEFAULT_PAGE_SIZE) int size) {
        return null;
    }
}
`

var constAppConstantsSource = `
package com.example.constants;

public class AppConstants {
    public static final String DEFAULT_PAGE_NUMBER = "0";
    public static final String DEFAULT_PAGE_SIZE = "30";
}
`
