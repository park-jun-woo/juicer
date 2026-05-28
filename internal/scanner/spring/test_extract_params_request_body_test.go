//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractParams_RequestBody — @RequestBody 추출 테스트
package spring

import "testing"

func TestExtractParams_RequestBody(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @PostMapping("/items")
    public void create(@RequestBody CreateItemRequest body) {}
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
	if ep.bodyType != "CreateItemRequest" {
		t.Errorf("body type: want CreateItemRequest, got %s", ep.bodyType)
	}
	if ep.bodyVarName != "body" {
		t.Errorf("body var name: want body, got %s", ep.bodyVarName)
	}
}
