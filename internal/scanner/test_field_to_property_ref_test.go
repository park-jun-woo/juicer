//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_Ref — Ref가 설정된 필드는 $ref로 변환된다
package scanner

import "testing"

func TestFieldToProperty_Ref(t *testing.T) {
	// Scalar named ref.
	p := fieldToProperty(Field{Name: "owner", Ref: "AlbumUserResponseDto"})
	if p["$ref"] != "#/components/schemas/AlbumUserResponseDto" {
		t.Fatalf("scalar ref: got %v", p)
	}

	// Array of named refs (Type=="array").
	p = fieldToProperty(Field{Name: "users", Type: "array", Ref: "AlbumUserResponseDto"})
	if p["type"] != "array" {
		t.Fatalf("array ref type: got %v", p)
	}
	items, ok := p["items"].(map[string]any)
	if !ok || items["$ref"] != "#/components/schemas/AlbumUserResponseDto" {
		t.Fatalf("array ref items: got %v", p)
	}
}
