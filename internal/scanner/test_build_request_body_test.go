//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildRequestBody 테스트
package scanner

import (
	"testing"
)

func TestBuildRequestBody(t *testing.T) {
	t.Run("file upload", func(t *testing.T) {
		req := &Request{
			Files:      []Param{{Name: "avatar", Type: "file"}},
			FormFields: []Param{{Name: "name", Type: "string"}},
		}
		schemas := map[string]any{}
		rb := buildRequestBody(req, schemas)
		if rb == nil {
			t.Fatal("expected requestBody")
		}
		content := rb["content"].(map[string]any)
		if content["multipart/form-data"] == nil {
			t.Error("expected multipart/form-data")
		}
	})

	t.Run("json body", func(t *testing.T) {
		req := &Request{
			Body: &Body{
				TypeName: "Req",
				Fields:   []Field{{Name: "name", Type: "string"}},
			},
		}
		schemas := map[string]any{}
		rb := buildRequestBody(req, schemas)
		if rb == nil {
			t.Fatal("expected requestBody")
		}
		content := rb["content"].(map[string]any)
		if content["application/json"] == nil {
			t.Error("expected application/json")
		}
	})

	t.Run("raw body", func(t *testing.T) {
		req := &Request{RawBody: true}
		schemas := map[string]any{}
		rb := buildRequestBody(req, schemas)
		if rb == nil {
			t.Fatal("expected requestBody")
		}
		content := rb["content"].(map[string]any)
		if content["application/octet-stream"] == nil {
			t.Error("expected application/octet-stream")
		}
	})

	t.Run("nil", func(t *testing.T) {
		req := &Request{}
		schemas := map[string]any{}
		rb := buildRequestBody(req, schemas)
		if rb != nil {
			t.Error("expected nil requestBody")
		}
	})
}
