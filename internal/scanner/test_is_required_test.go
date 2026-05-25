//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestIsRequired 테스트
package scanner

import (
	"testing"
)

func TestIsRequired(t *testing.T) {
	tests := []struct {
		validate string
		want     bool
	}{
		{"required", true},
		{"required,email", true},
		{"email,required", true},
		{"email", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.validate, func(t *testing.T) {
			f := Field{Validate: tt.validate}
			got := isRequired(f)
			if got != tt.want {
				t.Errorf("isRequired(%q) = %v, want %v", tt.validate, got, tt.want)
			}
		})
	}
}
