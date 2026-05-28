//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestDetectAuth_ClassLevel — 클래스 레벨 @PreAuthorize 테스트
package spring

import "testing"

func TestDetectAuth_ClassLevel(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
@PreAuthorize("hasRole('ADMIN')")
public class AdminController {

    @GetMapping("/dashboard")
    public void dashboard() {}
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fi := &fileInfo{root: root, src: src, relPath: "AdminController.java", absPath: "/test/AdminController.java"}
	fi.imports = extractImports(root, src)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	if len(controllers[0].roles) != 1 {
		t.Fatalf("expected 1 class-level role, got %d", len(controllers[0].roles))
	}
	if controllers[0].roles[0] != "ADMIN" {
		t.Errorf("class role: want ADMIN, got %s", controllers[0].roles[0])
	}
}
