//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what URL 경로에서 path parameter 추출 테스트
package laravel

import "testing"

func TestExtractURLParams(t *testing.T) {
	params := extractURLParams("/users/{user}/posts/{post}")
	if len(params) != 2 {
		t.Fatalf("expected 2 params, got %d", len(params))
	}
	if params[0].Name != "user" {
		t.Errorf("first param name = %q, want %q", params[0].Name, "user")
	}
	if params[1].Name != "post" {
		t.Errorf("second param name = %q, want %q", params[1].Name, "post")
	}
}

func TestExtractURLParams_None(t *testing.T) {
	params := extractURLParams("/users")
	if len(params) != 0 {
		t.Errorf("expected 0 params, got %d", len(params))
	}
}
