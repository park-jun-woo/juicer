//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what extractGenericInner — 제네릭 안쪽 타입 추출의 분기를 검증
package actix

import "testing"

func TestExtractGenericInner(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		// no '<' -> returned as-is
		{"String", "String"},
		// standard generic -> inner with trailing '>' stripped + trimmed
		{"Json<User>", "User"},
		{"web::Json< User >", "User"},
		// '<' present but no trailing '>' (malformed) -> inner kept, trimmed
		{"Json<User", "User"},
	}
	for _, c := range cases {
		if got := extractGenericInner(c.in); got != c.want {
			t.Errorf("extractGenericInner(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
