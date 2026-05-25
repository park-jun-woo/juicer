package scanner

import "testing"

func TestPathParams_WithParams(t *testing.T) {
	params := pathParams("/api/users/:id/posts/:postId")
	if len(params) != 2 {
		t.Fatalf("expected 2, got %d", len(params))
	}
	if params[0].Name != "id" {
		t.Fatalf("expected id, got %s", params[0].Name)
	}
}

func TestPathParams_NoParams(t *testing.T) {
	params := pathParams("/api/users")
	if len(params) != 0 {
		t.Fatalf("expected 0, got %d", len(params))
	}
}
