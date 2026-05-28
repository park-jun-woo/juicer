//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractParams_RequestHeader — @RequestHeader 추출 테스트
package spring

import "testing"

func TestExtractParams_RequestHeader(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @GetMapping("/items")
    public void list(@RequestHeader("X-Auth-Token") String token, @RequestHeader String accept) {}
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
	if len(ep.headers) != 2 {
		t.Fatalf("expected 2 header params, got %d", len(ep.headers))
	}
	if ep.headers[0].Name != "X-Auth-Token" {
		t.Errorf("headers[0] name: want X-Auth-Token, got %s", ep.headers[0].Name)
	}
	if ep.headers[0].Type != "string" {
		t.Errorf("headers[0] type: want string, got %s", ep.headers[0].Type)
	}
	if ep.headers[1].Name != "accept" {
		t.Errorf("headers[1] name: want accept, got %s", ep.headers[1].Name)
	}
}
