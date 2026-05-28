//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what PHP 타입 → OpenAPI 타입 변환 테스트
package laravel

import "testing"

func TestPhpTypeToOpenAPI(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"int", "integer"},
		{"integer", "integer"},
		{"float", "number"},
		{"string", "string"},
		{"bool", "boolean"},
		{"array", "array"},
		{"SomeClass", ""},
	}
	for _, tt := range tests {
		got := phpTypeToOpenAPI(tt.in)
		if got != tt.want {
			t.Errorf("phpTypeToOpenAPI(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
