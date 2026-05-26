//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what tsTypeToOpenAPIType 테스트
package nestjs

import "testing"

func TestTsTypeToOpenAPIType_Cases(t *testing.T) {
	cases := []struct{ in, want string }{
		{"string", "string"},
		{"number", "number"},
		{"boolean", "boolean"},
		{"Uuid", "string:uuid"},
		{"Date", "string:date-time"},
		{"MyDto", "object"},
		{"", "string"},
		{"void", "string"},
	}
	for _, c := range cases {
		got := tsTypeToOpenAPIType(c.in)
		if got != c.want {
			t.Errorf("tsTypeToOpenAPIType(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
