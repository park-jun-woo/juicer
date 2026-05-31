//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what quoteIdent 큰따옴표 인용 및 내부 따옴표 이스케이프 테스트
package prisma

import "testing"

func TestQuoteIdent(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"users", `"users"`},
		{"", `""`},
		{`we"ird`, `"we""ird"`},
		{`a"b"c`, `"a""b""c"`},
	}
	for _, c := range cases {
		if got := quoteIdent(c.in); got != c.want {
			t.Errorf("quoteIdent(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
