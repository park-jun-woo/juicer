//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what isNullableType 테스트
package fastapi

import "testing"

func TestIsNullableType(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"Optional[str]", true},
		{"str | None", true},
		{"None | str", true},
		{"Union[str, None]", true},
		{"str", false},
		{"int", false},
	}
	for _, tt := range tests {
		if got := isNullableType(tt.in); got != tt.want {
			t.Errorf("isNullableType(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
