//ff:func feature=scan type=convert control=sequence
//ff:what TestBuildResponses 테스트
package scanner

import (
	"testing"
)

func TestBuildResponses(t *testing.T) {
	t.Run("empty responses", func(t *testing.T) {
		schemas := map[string]any{}
		r := buildResponses(nil, schemas)
		if r["200"] == nil {
			t.Error("expected default 200 response")
		}
	})

	t.Run("json response with fields", func(t *testing.T) {
		schemas := map[string]any{}
		r := buildResponses([]Response{
			{Status: "200", Kind: "json", TypeName: "User", Fields: []Field{{Name: "name", Type: "string"}}},
		}, schemas)
		if r["200"] == nil {
			t.Error("expected 200 response")
		}
	})

	t.Run("partial confidence", func(t *testing.T) {
		schemas := map[string]any{}
		r := buildResponses([]Response{
			{Status: "200", Kind: "json", TypeName: "gin.H", Fields: []Field{{Name: "msg", Type: "string"}}, Confidence: "partial"},
		}, schemas)
		if r["200"] == nil {
			t.Error("expected 200 response")
		}
	})

	t.Run("non-json response", func(t *testing.T) {
		schemas := map[string]any{}
		r := buildResponses([]Response{
			{Status: "200", Kind: "string"},
		}, schemas)
		if r["200"] == nil {
			t.Error("expected 200 response")
		}
	})
}
