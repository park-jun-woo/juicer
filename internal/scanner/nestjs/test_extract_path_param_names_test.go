//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractPathParamNames 테스트
package nestjs

import "testing"

func TestExtractPathParamNames(t *testing.T) {
	if got := extractPathParamNames("/users/:id/posts/:postId"); len(got) != 2 || got[0] != "id" {
		t.Fatalf("got %v", got)
	}
	if got := extractPathParamNames("/users/{id}"); len(got) != 1 || got[0] != "id" {
		t.Fatalf("got %v", got)
	}
	if got := extractPathParamNames("/static/path"); got != nil {
		t.Fatalf("expected nil, got %v", got)
	}
}
