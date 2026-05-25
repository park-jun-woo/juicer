package scanner

import "testing"

func TestBuildRequestBody_JSON(t *testing.T) {
	req := &Request{Body: &Body{TypeName: "User", Fields: []Field{{Name: "name", JSON: "name"}}}}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBuildRequestBody_Multipart(t *testing.T) {
	req := &Request{Files: []Param{{Name: "file"}}}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBuildRequestBody_NoBody(t *testing.T) {
	req := &Request{}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result != nil {
		t.Fatal("expected nil")
	}
}
