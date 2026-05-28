//ff:func feature=scan type=test control=sequence topic=actix
//ff:what TestRustTypeToOpenAPI — Rust 타입 → OpenAPI 변환 테스트
package actix

import "testing"

func TestRustTypeToOpenAPI(t *testing.T) {
	tests := []struct {
		input      string
		wantType   string
		wantFormat string
		wantItems  string
	}{
		{"String", "string", "", ""},
		{"&str", "string", "", ""},
		{"i32", "integer", "int32", ""},
		{"i64", "integer", "int64", ""},
		{"u32", "integer", "int32", ""},
		{"u64", "integer", "int64", ""},
		{"f32", "number", "float", ""},
		{"f64", "number", "double", ""},
		{"bool", "boolean", "", ""},
		{"Uuid", "string", "uuid", ""},
		{"NaiveDateTime", "string", "date-time", ""},
		{"NaiveDate", "string", "date", ""},
		{"Vec<String>", "array", "", "String"},
		{"Option<i32>", "integer", "int32", ""},
		{"HashMap<String, String>", "object", "", ""},
		{"MyStruct", "object", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := rustTypeToOpenAPI(tt.input)
			if got.Type != tt.wantType {
				t.Errorf("type: want %s, got %s", tt.wantType, got.Type)
			}
			if got.Format != tt.wantFormat {
				t.Errorf("format: want %s, got %s", tt.wantFormat, got.Format)
			}
			if got.Items != tt.wantItems {
				t.Errorf("items: want %s, got %s", tt.wantItems, got.Items)
			}
		})
	}
}
