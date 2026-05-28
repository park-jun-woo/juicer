//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractParams_RequestParam — @RequestParam 추출 테스트
package spring

import "testing"

func TestExtractParams_RequestParam(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @GetMapping("/items")
    public void list(@RequestParam(defaultValue = "0") int page, @RequestParam String keyword) {}
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
	if len(ep.query) != 2 {
		t.Fatalf("expected 2 query params, got %d", len(ep.query))
	}
	if ep.query[0].Name != "page" {
		t.Errorf("query[0] name: want page, got %s", ep.query[0].Name)
	}
	if ep.query[0].Default != "0" {
		t.Errorf("query[0] default: want 0, got %s", ep.query[0].Default)
	}
	if ep.query[1].Name != "keyword" {
		t.Errorf("query[1] name: want keyword, got %s", ep.query[1].Name)
	}
}
