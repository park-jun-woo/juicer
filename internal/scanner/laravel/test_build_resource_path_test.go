//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestBuildResourcePath 테스트
package laravel

import "testing"

func TestBuildResourcePath(t *testing.T) {
	p, param := buildResourcePath("users")
	if p != "users" || param != "user" {
		t.Fatalf("single: %q %q", p, param)
	}
	p2, param2 := buildResourcePath("users.posts")
	if p2 != "users/{user}/posts" {
		t.Fatalf("nested path: %q", p2)
	}
	if param2 != "post" {
		t.Fatalf("nested param: %q", param2)
	}
}
