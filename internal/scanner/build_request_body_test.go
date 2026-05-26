//ff:func feature=scan type=test control=sequence
//ff:what TestBuildRequestBody_JSON 테스트
package scanner

import "testing"

func TestBuildRequestBody_JSON(t *testing.T) {
	req := &Request{Body: &Body{TypeName: "User", Fields: []Field{{Name: "name", JSON: "name"}}}}
	schemas := map[string]any{}
	result := buildRequestBody(req, schemas)
	if result == nil {
		t.Fatal("expected non-nil")
	}

	// multipart (files)
	reqFiles := &Request{Files: []Param{{Name: "file", Type: "file"}}}
	result = buildRequestBody(reqFiles, schemas)
	if result == nil {
		t.Fatal("expected non-nil for files")
	}

	// raw body
	reqRaw := &Request{RawBody: true}
	result = buildRequestBody(reqRaw, schemas)
	if result == nil {
		t.Fatal("expected non-nil for raw body")
	}

	// nil return (no body, no files, no rawBody)
	reqEmpty := &Request{}
	result = buildRequestBody(reqEmpty, schemas)
	if result != nil {
		t.Fatal("expected nil for empty request")
	}
}
