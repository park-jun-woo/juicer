//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestGoTypeToOpenAPI_DurationUUID time.Duration과 uuid.UUID 변환 반복 테스트
package scanner

import "testing"

func TestGoTypeToOpenAPI_DurationUUID(t *testing.T) {
	cases := []struct{ in, out string }{
		{"time.Duration", "string"},
		{"uuid.UUID", "string"},
	}
	for _, c := range cases {
		if got := goTypeToOpenAPI(c.in); got != c.out {
			t.Fatalf("goTypeToOpenAPI(%q)=%q, want %q", c.in, got, c.out)
		}
	}
}
