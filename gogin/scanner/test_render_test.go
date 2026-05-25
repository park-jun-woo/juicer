//ff:func feature=scan type=render control=sequence
//ff:what TestRender 테스트
package scanner

import (
	"testing"
)

func TestRender(t *testing.T) {
	result := &ScanResult{
		Endpoints: []Endpoint{
			{Method: "GET", Path: "/test"},
		},
	}

	t.Run("YAML", func(t *testing.T) {
		out, err := Render(result, FormatYAML)
		if err != nil {
			t.Fatalf("Render YAML error: %v", err)
		}
		if len(out) == 0 {
			t.Error("expected non-empty YAML output")
		}
	})

	t.Run("JSON", func(t *testing.T) {
		out, err := Render(result, FormatJSON)
		if err != nil {
			t.Fatalf("Render JSON error: %v", err)
		}
		if len(out) == 0 {
			t.Error("expected non-empty JSON output")
		}
	})

	t.Run("OpenAPI", func(t *testing.T) {
		out, err := Render(result, FormatOpenAPI)
		if err != nil {
			t.Fatalf("Render OpenAPI error: %v", err)
		}
		if len(out) == 0 {
			t.Error("expected non-empty OpenAPI output")
		}
	})

	t.Run("unknown format", func(t *testing.T) {
		_, err := Render(result, Format(99))
		if err == nil {
			t.Error("expected error for unknown format")
		}
	})
}
