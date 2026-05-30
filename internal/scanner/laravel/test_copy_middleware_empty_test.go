//ff:func feature=scan type=test control=sequence topic=laravel
//ff:what TestCopyMiddleware_Empty 테스트
package laravel

import "testing"

func TestCopyMiddleware_Empty(t *testing.T) {
	if cp := copyMiddleware(nil); len(cp) != 0 {
		t.Fatalf("got %v", cp)
	}
}
