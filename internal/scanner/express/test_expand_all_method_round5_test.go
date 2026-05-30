//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestExpandAllMethod_Round5 테스트
package express

import "testing"

func TestExpandAllMethod_Round5(t *testing.T) {
	if got := expandAllMethod("all"); len(got) != 5 {
		t.Fatalf("all: %v", got)
	}
	if got := expandAllMethod("GET"); len(got) != 1 || got[0] != "GET" {
		t.Fatalf("single: %v", got)
	}
}
