//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestGoTypeToOpenAPI 테스트
package scanner

import (
	"testing"
)

func TestGoTypeToOpenAPI(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"string", "string"},
		{"int", "integer"},
		{"int64", "integer"},
		{"float64", "number"},
		{"bool", "boolean"},
		{"boolean", "boolean"},
		{"integer", "integer"},
		{"number", "number"},
		{"object", "object"},
		{"array", "array"},
		{"any", "object"},
		{"interface{}", "object"},
		{"[]string", "array"},
		{"time.Time", "string"},
		{"CustomType", "object"},
		{"*string", "string"},
		{"*int", "integer"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := goTypeToOpenAPI(tt.input)
			if got != tt.want {
				t.Errorf("goTypeToOpenAPI(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
