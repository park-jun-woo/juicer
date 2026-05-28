//ff:func feature=scan type=test control=sequence topic=hono
//ff:what 경로에서 :param 추출 테스트
package hono

import "testing"

func TestExtractPathParams(t *testing.T) {
	params := extractPathParams("/users/:id/posts/:postId")
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	if params[0] != "id" {
		t.Errorf("expected id, got %s", params[0])
	}
	if params[1] != "postId" {
		t.Errorf("expected postId, got %s", params[1])
	}
}
