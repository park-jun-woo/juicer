//ff:type feature=scan type=test topic=spring
//ff:what E2E 테스트용 컨트롤러 소스 템플릿
package spring

var e2eControllerSource = `
package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.http.HttpStatus;
import org.springframework.security.access.prepost.PreAuthorize;
import com.example.demo.dto.CreateItemRequest;
import com.example.demo.dto.ItemDto;
import java.util.List;

@RestController
@RequestMapping("/api/items")
public class ItemController {

    @GetMapping
    public List<ItemDto> list(@RequestParam(defaultValue = "0") int page) {
        return null;
    }

    @GetMapping("/{id}")
    public ItemDto getById(@PathVariable Long id) {
        return null;
    }

    @PostMapping
    @ResponseStatus(HttpStatus.CREATED)
    @PreAuthorize("hasRole('ADMIN')")
    public ItemDto create(@RequestBody CreateItemRequest body) {
        return null;
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void delete(@PathVariable Long id) {}
}
`

var e2eReqDtoSource = `
package com.example.demo.dto;

import javax.validation.constraints.*;
import com.fasterxml.jackson.annotation.JsonProperty;

public class CreateItemRequest {
    @NotBlank
    @Size(min = 1, max = 100)
    private String name;

    @JsonProperty("item_price")
    @Min(0)
    private int price;

    private String description;
}
`

var e2eRespDtoSource = `
package com.example.demo.dto;

public class ItemDto {
    private Long id;
    private String name;
    private int price;
}
`
