//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what isPydanticModelType 테스트
package fastapi

import "testing"

func TestIsPydanticModelType(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"UserCreate", true},
		{"str", false},
		{"int", false},
		{"", false},
		{"list", false},
		{"Optional", false},
		{"MyModel", true},
	}
	for _, tt := range tests {
		if got := isPydanticModelType(tt.in); got != tt.want {
			t.Errorf("isPydanticModelType(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}
