//ff:func feature=scan type=extract control=sequence
//ff:what TestPathParams_WithParams 테스트
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
