package scanner

import "testing"

func TestBuildResponses_Empty(t *testing.T) {
	schemas := map[string]any{}
	result := buildResponses(nil, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBuildResponses_WithResponses(t *testing.T) {
	resps := []Response{{Status: "200"}, {Status: "400"}}
	schemas := map[string]any{}
	result := buildResponses(resps, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
