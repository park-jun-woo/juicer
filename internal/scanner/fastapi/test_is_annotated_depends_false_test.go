//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestIsAnnotatedDepends_False 테스트
package fastapi

import "testing"

func TestIsAnnotatedDepends_False(t *testing.T) {
	if isAnnotatedDepends("int", nil) {
		t.Fatal("expected false")
	}
	if isAnnotatedDepends("Annotated[int, Query()]", map[string]string{"x": "y"}) {
		t.Fatal("expected false for Annotated without Depends")
	}
}
