//ff:func feature=scan type=test control=sequence topic=spring
//ff:what Java 파싱 기본 테스트
package spring

import "testing"

func TestParseJava_Basic(t *testing.T) {
	src := []byte(`
package com.example.demo;

import org.springframework.web.bind.annotation.*;

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
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if root == nil {
		t.Fatal("root is nil")
	}
	if root.Type() != "program" {
		t.Errorf("root type: want program, got %s", root.Type())
	}
}
