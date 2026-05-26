//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what defaultStatusForMethod 테스트
package fastapi

import "testing"

func TestDefaultStatusForMethod(t *testing.T) {
	if got := defaultStatusForMethod("POST"); got != "201" {
		t.Errorf("POST: got %s, want 201", got)
	}
	if got := defaultStatusForMethod("GET"); got != "200" {
		t.Errorf("GET: got %s, want 200", got)
	}
	if got := defaultStatusForMethod("DELETE"); got != "200" {
		t.Errorf("DELETE: got %s, want 200", got)
	}
}
