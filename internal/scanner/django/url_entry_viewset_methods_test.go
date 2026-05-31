//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what urlEntryViewSetMethods — dict 우선·부모 역산 fallback을 검증
package django

import "testing"

func TestURLEntryViewSetMethods(t *testing.T) {
	vs := &viewsetInfo{name: "V", parents: []string{"ModelViewSet"}}

	// With an explicit dict, methods come straight from the dict.
	entry := urlEntry{methodActions: map[string]string{"get": "list", "post": "create"}}
	got := urlEntryViewSetMethods(entry, vs)
	if len(got) != 2 {
		t.Fatalf("expected 2 methods from dict, got %v", got)
	}
	byMethod := map[string]string{}
	for _, m := range got {
		byMethod[m.method] = m.action
	}
	if byMethod["GET"] != "list" || byMethod["POST"] != "create" {
		t.Errorf("dict methods mismatch: %v", byMethod)
	}

	// Without a dict, fall back to parent inference (ModelViewSet => 6 methods).
	fallback := urlEntryViewSetMethods(urlEntry{}, vs)
	if len(fallback) != 6 {
		t.Errorf("expected 6 inferred methods, got %d", len(fallback))
	}
}
