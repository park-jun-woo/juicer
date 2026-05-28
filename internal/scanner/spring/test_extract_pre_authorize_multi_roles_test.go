//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestExtractPreAuthorizeRoles_MultiHasRole — hasRole 다중 매칭 테스트
package spring

import "testing"

func TestExtractPreAuthorizeRoles_MultiHasRole(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class TestController {

    @GetMapping("/multi")
    @PreAuthorize("hasRole('ADMIN') or hasRole('MANAGER')")
    public void multi() {}
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
	if len(ep.roles) != 2 {
		t.Fatalf("expected 2 roles, got %d: %v", len(ep.roles), ep.roles)
	}
	want := map[string]bool{"ADMIN": true, "MANAGER": true}
	for _, r := range ep.roles {
		if !want[r] {
			t.Errorf("unexpected role: %s", r)
		}
	}
}
