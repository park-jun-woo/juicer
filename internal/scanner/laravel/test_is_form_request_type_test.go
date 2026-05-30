//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what TestIsFormRequestType 테스트
package laravel

import "testing"

func TestIsFormRequestType(t *testing.T) {
	if isFormRequestType("StoreUserRequest") != true {
		t.Fatal("expected true")
	}
	for _, n := range []string{"", "Request", "int", "string", "float", "bool", "array", "User"} {
		if isFormRequestType(n) {
			t.Errorf("expected false for %q", n)
		}
	}
}
