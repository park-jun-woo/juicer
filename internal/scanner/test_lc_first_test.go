//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestLcFirst 테스트
package scanner

import (
	"testing"
)

func TestLcFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Building", "building"},
		{"SMSResult", "smsResult"},
		{"ID", "id"},
		{"already", "already"},
		{"", ""},
		{"A", "a"},
		{"ABC", "abc"},
		{"ABCDef", "abcDef"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := lcFirst(tt.input)
			if got != tt.want {
				t.Errorf("lcFirst(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
