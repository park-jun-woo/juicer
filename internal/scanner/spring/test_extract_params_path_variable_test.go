//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractParams_PathVariable — @PathVariable 추출 테스트
package spring

import "testing"

func TestExtractParams_PathVariable(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @GetMapping("/items/{id}")
    public void get(@PathVariable("id") Long id) {}
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
	if len(ep.params) != 1 {
		t.Fatalf("expected 1 path param, got %d", len(ep.params))
	}
	if ep.params[0].Name != "id" {
		t.Errorf("param name: want id, got %s", ep.params[0].Name)
	}
}
