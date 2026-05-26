//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestGoTypeToOpenAPI_AllTypesCov 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_AllTypesCov(t *testing.T) {
	tests := []struct{ in, out string }{
		{"*string", "string"},
		{"int", "integer"}, {"int64", "integer"}, {"uint32", "integer"},
		{"float32", "number"}, {"float64", "number"},
		{"bool", "boolean"}, {"boolean", "boolean"},
		{"integer", "integer"}, {"number", "number"},
		{"object", "object"}, {"array", "array"},
		{"any", "object"}, {"interface{}", "object"},
		{"[]string", "array"},
		{"time.Time", "string"},
		{"SomeStruct", "object"},
	}
	for _, tt := range tests {
		got := goTypeToOpenAPI(tt.in)
		if got != tt.out {
			t.Errorf("goTypeToOpenAPI(%q) = %q, want %q", tt.in, got, tt.out)
		}
	}
}
