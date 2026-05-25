//ff:func feature=ratchet type=session control=iteration dimension=1
//ff:what TestToQueryName 테스트
package sqls

import (
	"testing"
)

func TestToQueryName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"UserRepository.GetAll", "UserGetAll"},
		{"UserRepo.GetAll", "UserRepoGetAll"},
		{"Simple", "Simple"},
		{"A.B", "AB"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := toQueryName(tt.input)
			if got != tt.want {
				t.Errorf("toQueryName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
