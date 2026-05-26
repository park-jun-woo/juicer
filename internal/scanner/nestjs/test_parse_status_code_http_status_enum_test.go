//ff:func feature=scan type=test control=iteration dimension=1 topic=nestjs
//ff:what TestParseStatusCode_HttpStatusEnum 테스트
package nestjs

import "testing"

func TestParseStatusCode_HttpStatusEnum(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"HttpStatus.OK", 200},
		{"HttpStatus.CREATED", 201},
		{"HttpStatus.NO_CONTENT", 204},
		{"HttpStatus.BAD_REQUEST", 400},
		{"HttpStatus.NOT_FOUND", 404},
		{"HttpStatus.INTERNAL_SERVER_ERROR", 500},
		{"HttpStatus.UNKNOWN", 0},
	}
	for _, c := range cases {
		got := parseStatusCode(c.in)
		if got != c.want {
			t.Errorf("parseStatusCode(%q) = %d, want %d", c.in, got, c.want)
		}
	}
}
