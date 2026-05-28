//ff:func feature=scan type=test control=iteration dimension=1 topic=spring
//ff:what TestExtractRoutes_MappingAnnotations — 매핑 어노테이션 추출 테스트
package spring

import "testing"

func TestExtractRoutes_MappingAnnotations(t *testing.T) {
	src := []byte(`
package com.example;

@RestController
@RequestMapping("/api/items")
public class ItemController {

    @GetMapping
    public void list() {}

    @PostMapping("/create")
    public void create() {}

    @PutMapping("/{id}")
    public void update() {}

    @DeleteMapping("/{id}")
    public void delete() {}

    @PatchMapping("/{id}")
    public void patch() {}
}
`)
	root, err := parseJava(src)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	fi := &fileInfo{root: root, src: src, relPath: "ItemController.java", absPath: "/test/ItemController.java"}
	fi.imports = extractImports(root, src)
	controllers := extractControllers(fi)
	if len(controllers) != 1 {
		t.Fatalf("expected 1 controller, got %d", len(controllers))
	}
	ci := controllers[0]
	if ci.prefix != "/api/items" {
		t.Errorf("prefix: want /api/items, got %s", ci.prefix)
	}
	if len(ci.endpoints) != 5 {
		t.Fatalf("expected 5 endpoints, got %d", len(ci.endpoints))
	}

	tests := []struct {
		method  string
		path    string
		handler string
	}{
		{"GET", "", "list"},
		{"POST", "/create", "create"},
		{"PUT", "/{id}", "update"},
		{"DELETE", "/{id}", "delete"},
		{"PATCH", "/{id}", "patch"},
	}
	for i, tt := range tests {
		ep := ci.endpoints[i]
		if ep.method != tt.method {
			t.Errorf("ep[%d] method: want %s, got %s", i, tt.method, ep.method)
		}
		if ep.path != tt.path {
			t.Errorf("ep[%d] path: want %s, got %s", i, tt.path, ep.path)
		}
		if ep.handler != tt.handler {
			t.Errorf("ep[%d] handler: want %s, got %s", i, tt.handler, ep.handler)
		}
	}
}
