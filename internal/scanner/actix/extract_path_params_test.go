//ff:func feature=scan type=test control=sequence topic=actix
//ff:what extractPathParams — {param} 경로 파라미터 추출을 검증
package actix

import "testing"

func TestExtractPathParams(t *testing.T) {
	// Mixed: literal segments, two params.
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

func TestExtractPathParams_None(t *testing.T) {
	// No braces -> HasPrefix/HasSuffix false branch.
	if p := extractPathParams("/users/list"); len(p) != 0 {
		t.Fatalf("expected no params, got %+v", p)
	}
}
