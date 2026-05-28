//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractTypeParams — 클래스 선언에서 타입 파라미터 추출 테스트
package spring

import "testing"

func TestExtractTypeParams(t *testing.T) {
	src := []byte(`
package com.example;

public class PagedResponse<T> {
    private List<T> content;
    private int totalPages;
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	classes := findAllByType(root, "class_declaration")
	if len(classes) == 0 {
		t.Fatal("no class found")
	}
	params := extractTypeParams(classes[0], src)
	if len(params) != 1 || params[0] != "T" {
		t.Errorf("extractTypeParams = %v, want [T]", params)
	}
}
