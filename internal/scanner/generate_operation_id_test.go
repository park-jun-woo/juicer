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

	// inline handler
	ep2 := Endpoint{Handler: "(inline)", Method: "GET", Path: "/api/health"}
	got2 := generateOperationID(ep2)
	if got2 == "" {
		t.Fatal("expected non-empty for inline")
	}

	// empty handler
	ep3 := Endpoint{Method: "POST", Path: "/users"}
	got3 := generateOperationID(ep3)
	if got3 == "" {
		t.Fatal("expected non-empty for empty handler")
	}
}

