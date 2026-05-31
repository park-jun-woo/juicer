//ff:func feature=prisma type=test control=iteration dimension=1 topic=prisma
//ff:what unsupportedType 내부 원문 추출/이스케이프 해제 테스트
package prisma

import "testing"

func TestUnsupportedType(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`Unsupported("vector(768)")`, "vector(768)"},
		{`Unsupported("\"public\".geography")`, `"public".geography`},
		{"NoQuotes", "NoQuotes"},
		{`Only"OneQuote`, `Only"OneQuote`},
	}
	for _, c := range cases {
		if got := unsupportedType(c.in); got != c.want {
			t.Errorf("unsupportedType(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
