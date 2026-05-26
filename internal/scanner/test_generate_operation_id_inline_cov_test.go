//ff:func feature=scan type=test control=sequence
//ff:what TestGenerateOperationID_InlineCov 테스트
package scanner

import "testing"

func TestGenerateOperationID_InlineCov(t *testing.T) {
	ep := Endpoint{Handler: "(inline)", Method: "GET", Path: "/api/health"}
	got := generateOperationID(ep)
	if got == "" {
		t.Fatal("expected non-empty operationId")
	}
}
