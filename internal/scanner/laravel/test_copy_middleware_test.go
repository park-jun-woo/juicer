//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCopyMiddleware 테스트
package laravel

import "testing"

func TestCopyMiddleware(t *testing.T) {
	orig := []string{"auth", "throttle"}
	cp := copyMiddleware(orig)
	if len(cp) != 2 || cp[0] != "auth" {
		t.Fatalf("got %v", cp)
	}
	cp[0] = "changed"
	if orig[0] != "auth" {
		t.Fatal("copy shares backing array")
	}
}
