//ff:func feature=scan type=test control=sequence
//ff:what hasAuthMiddleware 테스트
package scanner

import "testing"

func TestHasAuthMiddleware(t *testing.T) {
	if !hasAuthMiddleware([]string{"logger", "AuthRequired"}) {
		t.Error("expected true when an auth middleware is present")
	}
	if hasAuthMiddleware([]string{"logger", "cors"}) {
		t.Error("expected false for non-auth middlewares")
	}
	if hasAuthMiddleware(nil) {
		t.Error("expected false for empty list")
	}
}
