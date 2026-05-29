//ff:func feature=scan type=test control=sequence
//ff:what TestResponseContent 테스트
package scanner

import "testing"

func TestResponseContent(t *testing.T) {
	t.Run("empty kind yields nil", func(t *testing.T) {
		if responseContent(Response{Kind: "empty"}, map[string]any{}) != nil {
			t.Error("expected nil content for empty kind")
		}
	})

	t.Run("json with fields yields application/json", func(t *testing.T) {
		c := responseContent(Response{Kind: "json", Fields: []Field{{Name: "id", Type: "int"}}}, map[string]any{})
		if c["application/json"] == nil {
			t.Fatalf("expected application/json content, got %v", c)
		}
	})

	t.Run("text yields text/plain string schema", func(t *testing.T) {
		c := responseContent(Response{Kind: "text"}, map[string]any{})
		media, ok := c["text/plain"].(map[string]any)
		if !ok {
			t.Fatalf("expected text/plain content, got %v", c)
		}
		schema, ok := media["schema"].(map[string]any)
		if !ok || schema["type"] != "string" {
			t.Fatalf("expected string schema, got %v", media)
		}
	})

	t.Run("string kind yields text/plain", func(t *testing.T) {
		c := responseContent(Response{Kind: "string"}, map[string]any{})
		if c["text/plain"] == nil {
			t.Fatalf("expected text/plain content, got %v", c)
		}
	})

	t.Run("explicit content type honored", func(t *testing.T) {
		c := responseContent(Response{Kind: "json", ContentType: "application/xml", Fields: []Field{{Name: "id", Type: "int"}}}, map[string]any{})
		if c["application/xml"] == nil {
			t.Fatalf("expected application/xml content, got %v", c)
		}
	})
}
