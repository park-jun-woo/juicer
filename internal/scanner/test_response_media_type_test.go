//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestResponseMediaType 테스트
package scanner

import "testing"

func TestResponseMediaType(t *testing.T) {
	cases := []struct {
		resp Response
		want string
	}{
		{Response{ContentType: "application/xml"}, "application/xml"},
		{Response{Kind: "text"}, "text/plain"},
		{Response{Kind: "string"}, "text/plain"},
		{Response{Kind: "json"}, "application/json"},
		{Response{Kind: "empty"}, "application/json"},
		{Response{ContentType: "text/csv", Kind: "json"}, "text/csv"},
	}
	for _, c := range cases {
		if got := responseMediaType(c.resp); got != c.want {
			t.Errorf("responseMediaType(%+v) = %q, want %q", c.resp, got, c.want)
		}
	}
}
