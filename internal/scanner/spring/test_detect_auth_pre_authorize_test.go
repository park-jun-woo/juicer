//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestDetectAuth_PreAuthorize — @PreAuthorize 역할 추출 테스트
package spring

import "testing"

func TestDetectAuth_PreAuthorize(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @GetMapping("/admin")
    @PreAuthorize("hasRole('ADMIN')")
    public void admin() {}
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
	if len(ep.roles) != 1 {
		t.Fatalf("expected 1 role, got %d", len(ep.roles))
	}
	if ep.roles[0] != "ADMIN" {
		t.Errorf("role: want ADMIN, got %s", ep.roles[0])
	}
}
