//ff:func feature=scan type=test control=iteration dimension=1 topic=flask
//ff:what Flask URL 컨버터 타입을 OpenAPI 타입으로 변환한다
package flask

import "testing"

func TestFlaskConverterToOpenAPI(t *testing.T) {
	tests := []struct {
		converter string
		wantType  string
		wantFmt   string
	}{
		{"int", "integer", "int64"},
		{"float", "number", "double"},
		{"uuid", "string", "uuid"},
		{"path", "string", ""},
		{"string", "string", ""},
		{"", "string", ""},
		{"custom_converter", "string", ""}, // default branch
	}
	for _, tt := range tests {
		got := flaskConverterToOpenAPI(tt.converter)
		if got.Type != tt.wantType {
			t.Errorf("flaskConverterToOpenAPI(%q).Type = %q, want %q", tt.converter, got.Type, tt.wantType)
		}
		if got.Format != tt.wantFmt {
			t.Errorf("flaskConverterToOpenAPI(%q).Format = %q, want %q", tt.converter, got.Format, tt.wantFmt)
		}
	}
}
