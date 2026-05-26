//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what unwrapPromise 테스트
package nestjs

import "testing"

func TestUnwrapPromise_Cases(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Promise<string>", "string"},
		{"string", "string"},
		{"Promise<MyDto>", "MyDto"},
		{"", ""},
	}
	for _, c := range cases {
		got := unwrapPromise(c.in)
		if got != c.want {
			t.Errorf("unwrapPromise(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
