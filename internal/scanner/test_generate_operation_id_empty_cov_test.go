//ff:func feature=scan type=test control=sequence
//ff:what TestGenerateOperationID_EmptyCov 테스트
package scanner

import "testing"

func TestGenerateOperationID_EmptyCov(t *testing.T) {
	ep := Endpoint{Handler: "", Method: "POST", Path: "/api/users"}
	got := generateOperationID(ep)
	if got == "" {
		t.Fatal("expected non-empty operationId")
	}
}
