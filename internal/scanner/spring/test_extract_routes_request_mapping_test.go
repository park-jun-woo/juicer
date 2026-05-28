//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractRoutes_RequestMappingMethod — @RequestMapping method 속성 추출 테스트
package spring

import "testing"

func TestExtractRoutes_RequestMappingMethod(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
public class LegacyController {

    @RequestMapping(value = "/old", method = RequestMethod.POST)
    public void oldEndpoint() {}
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fi := &fileInfo{root: root, src: src, relPath: "LegacyController.java", absPath: "/test/LegacyController.java"}
	fi.imports = extractImports(root, src)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	if len(controllers[0].endpoints) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(controllers[0].endpoints))
	}
	ep := controllers[0].endpoints[0]
	if ep.method != "POST" {
		t.Errorf("method: want POST, got %s", ep.method)
	}
	if ep.path != "/old" {
		t.Errorf("path: want /old, got %s", ep.path)
	}
}
