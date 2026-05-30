//ff:func feature=scan type=test control=sequence topic=hono
//ff:what TestFindObjectValueByKey_NilObj 테스트
package hono

import "testing"

func TestFindObjectValueByKey_NilObj(t *testing.T) {
	if findObjectValueByKey(nil, "x", []byte("")) != nil {
		t.Fatal("expected nil for nil obj")
	}
}
