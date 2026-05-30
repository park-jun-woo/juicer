//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestExtractPathParams 테스트
package actix

import "testing"

func TestExtractPathParams(t *testing.T) {

	params := extractPathParams("/users/{id}/posts/{postId}")
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d: %+v", len(params), params)
	}
	if params[0].Name != "id" || params[0].Type != "string" {
		t.Errorf("params[0] = %+v, want {id string}", params[0])
	}
	if params[1].Name != "postId" {
		t.Errorf("params[1].Name = %q, want postId", params[1].Name)
	}
}
