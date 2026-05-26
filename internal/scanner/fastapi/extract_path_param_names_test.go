//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractPathParamNames 테스트
package fastapi

import "testing"

func TestExtractPathParamNames(t *testing.T) {
	names := extractPathParamNames("/users/{user_id}/posts/{post_id}")
	if !names["user_id"] || !names["post_id"] {
		t.Fatalf("unexpected: %v", names)
	}
	if len(names) != 2 {
		t.Fatalf("expected 2, got %d", len(names))
	}

	empty := extractPathParamNames("/users")
	if len(empty) != 0 {
		t.Fatalf("expected 0, got %d", len(empty))
	}

	// malformed
	partial := extractPathParamNames("/users/{id")
	if len(partial) != 0 {
		t.Fatalf("expected 0 for malformed, got %d", len(partial))
	}
}
