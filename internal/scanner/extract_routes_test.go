package scanner

import "testing"

func TestExtractRoutes_NilPkgs(t *testing.T) {
	result := extractRoutes(nil, ".")
	if len(result) != 0 {
		t.Fatalf("expected 0, got %d", len(result))
	}
}
