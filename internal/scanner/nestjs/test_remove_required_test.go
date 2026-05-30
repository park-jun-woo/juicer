//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestRemoveRequired 테스트
package nestjs

import "testing"

func TestRemoveRequired(t *testing.T) {
	if got := removeRequired("required,email"); got != "email" {
		t.Fatalf("got %q", got)
	}
	if got := removeRequired("required"); got != "" {
		t.Fatalf("got %q", got)
	}
	if got := removeRequired("email,min:1"); got != "email,min:1" {
		t.Fatalf("got %q", got)
	}
}
