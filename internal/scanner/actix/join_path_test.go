//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what joinPath — 경로 세그먼트 결합 분기를 검증
package actix

import "testing"

func TestJoinPath(t *testing.T) {
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"/api", "/users"}, "/api/users"},
		{[]string{"api/", "users"}, "/api/users"},
		{[]string{" /v1 ", "items"}, "/v1/items"},
		// all empty/slash segments -> result stays "/"
		{[]string{"", "/", "  "}, "/"},
		{[]string{}, "/"},
		{[]string{"/single/"}, "/single"},
	}
	for _, c := range cases {
		if got := joinPath(c.in...); got != c.want {
			t.Errorf("joinPath(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
