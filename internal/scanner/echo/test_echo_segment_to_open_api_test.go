//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestEchoSegmentToOpenAPI 테스트
package echo

import "testing"

func TestEchoSegmentToOpenAPI(t *testing.T) {
	cases := map[string]string{
		":id":   "{id}",
		"*name": "{name}",
		"*":     "{wildcard}",
		"books": "books",
	}
	for in, want := range cases {
		if got := echoSegmentToOpenAPI(in); got != want {
			t.Fatalf("%s: expected %s, got %s", in, want, got)
		}
	}
}
