//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractResponse_ReturnType — 반환 타입 추출 테스트
package spring

import "testing"

func TestExtractResponse_ReturnType(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @GetMapping("/items")
    public ItemDto get() { return null; }
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fi := &fileInfo{root: root, src: src, relPath: "TestController.java", absPath: "/test/TestController.java"}
	fi.imports = extractImports(root, src)
	controllers := extractControllers(fi)
	if len(controllers) != 1 || len(controllers[0].endpoints) != 1 {
		t.Fatalf("expected 1 controller with 1 endpoint")
	}
	ep := controllers[0].endpoints[0]
	if ep.returnType != "ItemDto" {
		t.Errorf("return type: want ItemDto, got %s", ep.returnType)
	}
}
