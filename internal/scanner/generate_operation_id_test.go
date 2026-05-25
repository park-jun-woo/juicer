package scanner

import "testing"

func TestGenerateOperationID_Handler(t *testing.T) {
	ep := Endpoint{Handler: "file.go:h.ListBuildings", Method: "GET", Path: "/buildings"}
	got := generateOperationID(ep)
	if got != "listBuildings" {
		t.Fatalf("expected listBuildings, got %s", got)
	}
}

func TestGenerateOperationID_Inline(t *testing.T) {
	ep := Endpoint{Handler: "(inline)", Method: "GET", Path: "/api/v1/users"}
	got := generateOperationID(ep)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGenerateOperationID_Empty(t *testing.T) {
	ep := Endpoint{Method: "POST", Path: "/api/v1/items"}
	got := generateOperationID(ep)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}

func TestGenerateOperationID_WithParens(t *testing.T) {
	ep := Endpoint{Handler: "h.Create()", Method: "POST", Path: "/items"}
	got := generateOperationID(ep)
	if got != "create" {
		t.Fatalf("expected create, got %s", got)
	}
}
