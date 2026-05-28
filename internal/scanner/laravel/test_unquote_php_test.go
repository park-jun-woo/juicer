//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what PHP 문자열 따옴표 제거 테스트
package laravel

import "testing"

func TestUnquotePHP(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{`'hello'`, "hello"},
		{`"hello"`, "hello"},
		{`hello`, "hello"},
		{`''`, ""},
		{`""`, ""},
		{`x`, "x"},
	}
	for _, tt := range tests {
		got := unquotePHP(tt.in)
		if got != tt.want {
			t.Errorf("unquotePHP(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
