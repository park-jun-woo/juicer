//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryPipeNullable 테스트
package fastapi

import "testing"

func TestTryPipeNullable(t *testing.T) {
	oa, ok := tryPipeNullable("str | None")
	if !ok || oa.Type != "string" || !oa.Nullable {
		t.Fatalf("got %v, %v", oa, ok)
	}

	_, ok = tryPipeNullable("str | int")
	if ok {
		t.Fatal("should not match non-None union")
	}

	_, ok = tryPipeNullable("str")
	if ok {
		t.Fatal("should not match single type")
	}
}
