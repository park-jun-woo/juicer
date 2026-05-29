//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
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
		input, wantType, wantFormat, wantItems := tt.input, tt.wantType, tt.wantFormat, tt.wantItems
		t.Run(input, func(t *testing.T) {
			assertRustType(t, input, wantType, wantFormat, wantItems)
		})
	}
}
