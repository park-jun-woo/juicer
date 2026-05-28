//ff:func feature=scan type=test control=sequence topic=spring
//ff:what DTO 필드 추출 테스트 — 필드명, 타입, 유효성 검증 어노테이션
package spring

import "testing"

func TestExtractClassFields_Basic(t *testing.T) {
	src := []byte(`
package com.example;

public class CreateUserRequest {
    @NotBlank
    private String name;

    @NotNull
    @JsonProperty("email_address")
    private String email;

    @Min(0)
    @Max(150)
    private int age;

    @Size(min = 2, max = 100)
    private String nickname;

    @Email
    private String contactEmail;
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) != 1 {
		t.Fatalf("expected 1 class, got %d", len(classes))
	}
	fields := extractClassFields(classes[0], src)
	if len(fields) != 5 {
		t.Fatalf("expected 5 fields, got %d", len(fields))
	}

	if fields[0].Name != "name" {
		t.Errorf("field[0] name: want name, got %s", fields[0].Name)
	}
	if fields[0].Validate != "required" {
		t.Errorf("field[0] validate: want required, got %s", fields[0].Validate)
	}

	if fields[1].Name != "email" {
		t.Errorf("field[1] name: want email, got %s", fields[1].Name)
	}
	if fields[1].JSON != "email_address" {
		t.Errorf("field[1] json: want email_address, got %s", fields[1].JSON)
	}
	if fields[1].Validate != "required" {
		t.Errorf("field[1] validate: want required, got %s", fields[1].Validate)
	}

	if fields[2].Name != "age" {
		t.Errorf("field[2] name: want age, got %s", fields[2].Name)
	}
	if fields[2].Minimum == nil || *fields[2].Minimum != 0 {
		t.Errorf("field[2] minimum: want 0, got %v", fields[2].Minimum)
	}
	if fields[2].Maximum == nil || *fields[2].Maximum != 150 {
		t.Errorf("field[2] maximum: want 150, got %v", fields[2].Maximum)
	}

	if fields[3].Name != "nickname" {
		t.Errorf("field[3] name: want nickname, got %s", fields[3].Name)
	}
	if fields[3].MinLength == nil || *fields[3].MinLength != 2 {
		t.Errorf("field[3] minLength: want 2, got %v", fields[3].MinLength)
	}
	if fields[3].MaxLength == nil || *fields[3].MaxLength != 100 {
		t.Errorf("field[3] maxLength: want 100, got %v", fields[3].MaxLength)
	}

	if fields[4].Name != "contactEmail" {
		t.Errorf("field[4] name: want contactEmail, got %s", fields[4].Name)
	}
	if fields[4].Type != "string:email" {
		t.Errorf("field[4] type: want string:email, got %s", fields[4].Type)
	}
}
