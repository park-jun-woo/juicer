//ff:func feature=scan type=extract control=sequence
//ff:what TestGenerateOperationID_Inline 테스트
package scanner

import "testing"

func TestGenerateOperationID_Inline(t *testing.T) {
	ep := Endpoint{Handler: "(inline)", Method: "GET", Path: "/api/v1/users"}
	got := generateOperationID(ep)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
