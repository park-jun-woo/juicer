//ff:func feature=scan type=convert control=iteration dimension=1
//ff:what TestStatusDescription 테스트
package scanner

import (
	"testing"
)

func TestStatusDescription(t *testing.T) {
	tests := []struct {
		status string
		want   string
	}{
		{"200", "OK"},
		{"404", "Not Found"},
		{"(unknown)", "Error"},
		{"418", "Response"},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			got := statusDescription(tt.status)
			if got != tt.want {
				t.Errorf("statusDescription(%q) = %q, want %q", tt.status, got, tt.want)
			}
		})
	}
}
