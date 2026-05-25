//ff:func feature=scan type=extract control=sequence
//ff:what TestGenerateOperationID_WithParens 테스트
package scanner

import "testing"

func TestGenerateOperationID_WithParens(t *testing.T) {
	ep := Endpoint{Handler: "h.Create()", Method: "POST", Path: "/items"}
	got := generateOperationID(ep)
	if got != "create" {
		t.Fatalf("expected create, got %s", got)
	}
}
