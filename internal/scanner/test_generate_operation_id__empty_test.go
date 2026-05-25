//ff:func feature=scan type=extract control=sequence
//ff:what TestGenerateOperationID_Empty 테스트
package scanner

import "testing"

func TestGenerateOperationID_Empty(t *testing.T) {
	ep := Endpoint{Method: "POST", Path: "/api/v1/items"}
	got := generateOperationID(ep)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
