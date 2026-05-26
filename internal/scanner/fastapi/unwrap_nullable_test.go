//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what unwrapNullable 테스트
package fastapi

import "testing"

func TestUnwrapNullable(t *testing.T) {
	tests := []struct{ in, want string }{
		{"Optional[str]", "str"},
		{"str | None", "str"},
		{"str", "str"},
		{"int", "int"},
	}
	for _, tt := range tests {
		got := unwrapNullable(tt.in)
		if got != tt.want {
			t.Errorf("unwrapNullable(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
