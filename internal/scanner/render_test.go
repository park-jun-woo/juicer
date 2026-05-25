package scanner

import "testing"

func TestRender_YAML(t *testing.T) {
	result := &ScanResult{}
	data, err := Render(result, FormatYAML)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestRender_JSON(t *testing.T) {
	result := &ScanResult{}
	data, err := Render(result, FormatJSON)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestRender_OpenAPI(t *testing.T) {
	result := &ScanResult{}
	data, err := Render(result, FormatOpenAPI)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Fatal("expected non-empty")
	}
}

func TestRender_Unknown(t *testing.T) {
	result := &ScanResult{}
	_, err := Render(result, Format(99))
	if err == nil {
		t.Fatal("expected error")
	}
}
