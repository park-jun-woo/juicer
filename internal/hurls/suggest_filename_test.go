package hurls

import "testing"

func TestSuggestFilename_Basic(t *testing.T) {
	got := suggestFilename("/api/v1/admin/buildings")
	if got != "buildings.hurl" {
		t.Fatalf("got %q", got)
	}
}

func TestSuggestFilename_WithParam(t *testing.T) {
	got := suggestFilename("/api/v1/admin/buildings/:id")
	if got != "buildings_id.hurl" {
		t.Fatalf("got %q", got)
	}
}

func TestSuggestFilename_Health(t *testing.T) {
	got := suggestFilename("/api/health")
	if got != "health.hurl" {
		t.Fatalf("got %q", got)
	}
}

func TestSuggestFilename_Empty(t *testing.T) {
	got := suggestFilename("/api/v1")
	if got != "test.hurl" {
		t.Fatalf("got %q", got)
	}
}
