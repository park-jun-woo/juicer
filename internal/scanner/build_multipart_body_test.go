package scanner

import "testing"

func TestBuildMultipartBody(t *testing.T) {
	req := &Request{
		FormFields: []Param{{Name: "title"}},
		Files:      []Param{{Name: "file"}},
	}
	result := buildMultipartBody(req)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}

func TestBuildMultipartBody_Empty(t *testing.T) {
	req := &Request{}
	result := buildMultipartBody(req)
	if result == nil {
		t.Fatal("expected non-nil")
	}
}
