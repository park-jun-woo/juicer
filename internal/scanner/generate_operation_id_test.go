//ff:func feature=scan type=test control=sequence
//ff:what TestGenerateOperationID_Handler 테스트
package scanner

import "testing"

func TestGenerateOperationID_Handler(t *testing.T) {
	ep := Endpoint{Handler: "file.go:h.ListBuildings", Method: "GET", Path: "/buildings"}
	got := generateOperationID(ep)
	if got != "listBuildings" {
		t.Fatalf("expected listBuildings, got %s", got)
	}
}

