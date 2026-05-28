//ff:type feature=scan type=test topic=spring
//ff:what 인터페이스 상속 엔드포인트 테스트용 소스 템플릿
package spring

var interfaceApiSource = `
package com.example.api;

import org.springframework.web.bind.annotation.*;
import org.springframework.http.ResponseEntity;

@RequestMapping("/api/owners")
public interface OwnersApi {

    @GetMapping("/{ownerId}")
    ResponseEntity<OwnerDto> getOwner(@PathVariable int ownerId);

    @PostMapping
    ResponseEntity<OwnerDto> createOwner(@RequestBody CreateOwnerRequest body);
}
`

var interfaceControllerSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;
import com.example.api.OwnersApi;

@RestController
public class OwnerRestController implements OwnersApi {

    @Override
    public ResponseEntity<OwnerDto> getOwner(int ownerId) {
        return null;
    }

    @Override
    public ResponseEntity<OwnerDto> createOwner(CreateOwnerRequest body) {
        return null;
    }
}
`

var interfaceControllerWithPrefixSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;
import com.example.api.OwnersApi;

@RestController
@RequestMapping("/v2/owners")
public class OwnerRestControllerV2 implements OwnersApi {

    @Override
    public ResponseEntity<OwnerDto> getOwner(int ownerId) {
        return null;
    }

    @Override
    public ResponseEntity<OwnerDto> createOwner(CreateOwnerRequest body) {
        return null;
    }
}
`

var directAnnotationControllerSource = `
package com.example.controller;

import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/pets")
public class PetController {

    @GetMapping("/{petId}")
    public String getPet(@PathVariable int petId) {
        return null;
    }

    @PostMapping
    public String createPet(@RequestBody String body) {
        return null;
    }
}
`
