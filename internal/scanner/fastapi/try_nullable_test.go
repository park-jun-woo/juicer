//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what tryNullable 테스트
package fastapi

import "testing"

func TestTryNullable(t *testing.T) {
	// Optional
	oa, ok := tryNullable("Optional[str]")
	if !ok || oa.Type != "string" || !oa.Nullable {
		t.Fatalf("Optional: got %v, %v", oa, ok)
	}
	// Union
	oa2, ok := tryNullable("Union[int, None]")
	if !ok || oa2.Type != "integer" || !oa2.Nullable {
		t.Fatalf("Union: got %v, %v", oa2, ok)
	}
	// Pipe
	oa3, ok := tryNullable("float | None")
	if !ok || oa3.Type != "number" || !oa3.Nullable {
		t.Fatalf("Pipe: got %v, %v", oa3, ok)
	}
	// Plain
	_, ok = tryNullable("str")
	if ok {
		t.Fatal("should not match plain type")
	}
}
