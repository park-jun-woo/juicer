//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractResponse_ResponseStatus — @ResponseStatus 추출 테스트
package spring

import "testing"

func TestExtractResponse_ResponseStatus(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @PostMapping("/items")
    @ResponseStatus(HttpStatus.CREATED)
    public void create() {}
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
	if ep.statusCode != "201" {
		t.Errorf("status code: want 201, got %s", ep.statusCode)
	}
}
