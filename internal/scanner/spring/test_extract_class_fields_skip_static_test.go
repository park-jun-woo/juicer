//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestExtractClassFields_SkipStatic — static 필드가 제외되는지 확인
package spring

import "testing"

func TestExtractClassFields_SkipStatic(t *testing.T) {
	src := []byte(`
package com.example;

public class UserDto {
    private static final long serialVersionUID = 1L;
    private String name;
    private int age;
    private static String CONSTANT = "foo";
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("class not found")
	}
	cls := classes[0]
	fields := extractClassFields(cls, src)
	for _, f := range fields {
		if f.Name == "serialVersionUID" || f.Name == "CONSTANT" {
			t.Errorf("static field %q should have been skipped", f.Name)
		}
	}
	if len(fields) != 2 {
		t.Errorf("expected 2 fields (name, age), got %d: %+v", len(fields), fields)
	}
}
