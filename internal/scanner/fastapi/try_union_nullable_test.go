//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryUnionNullable 테스트
package fastapi

import "testing"

func TestTryUnionNullable(t *testing.T) {
	oa, ok := tryUnionNullable("Union[str, None]")
	if !ok || oa.Type != "string" || !oa.Nullable {
		t.Fatalf("got %v, %v", oa, ok)
	}

	_, ok = tryUnionNullable("Union[str, int]")
	if ok {
		t.Fatal("should not match non-None union")
	}
}
