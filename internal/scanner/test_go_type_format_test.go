//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestGoTypeFormat 테스트
package scanner

import (
	"testing"
)

func TestGoTypeFormat(t *testing.T) {
	tests := []struct {
		goType   string
		validate string
		want     string
	}{
		{"int64", "", "int64"},
		{"uint64", "", "int64"},
		{"int32", "", "int32"},
		{"uint32", "", "int32"},
		{"float64", "", "double"},
		{"float32", "", "float"},
		{"time.Time", "", "date-time"},
		{"string", "email", "email"},
		{"string", "url", "uri"},
		{"string", "uri", "uri"},
		{"string", "", ""},
		{"int", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.goType+"_"+tt.validate, func(t *testing.T) {
			f := Field{Validate: tt.validate}
			got := goTypeFormat(tt.goType, f)
			if got != tt.want {
				t.Errorf("goTypeFormat(%q) = %q, want %q", tt.goType, got, tt.want)
			}
		})
	}
}
