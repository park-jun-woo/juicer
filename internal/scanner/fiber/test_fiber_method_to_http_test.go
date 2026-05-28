//ff:func feature=scan type=test control=iteration dimension=1
//ff:what TestFiberMethodToHTTP 테스트
package fiber

import "testing"

func TestFiberMethodToHTTP(t *testing.T) {
	cases := map[string]string{
		"Get":     "GET",
		"Post":    "POST",
		"Put":     "PUT",
		"Patch":   "PATCH",
		"Delete":  "DELETE",
		"Head":    "HEAD",
		"Options": "OPTIONS",
		"All":     "ALL",
	}
	for method, expected := range cases {
		got := fiberMethodToHTTP[method]
		if got != expected {
			t.Errorf("fiberMethodToHTTP[%s] = %s, want %s", method, got, expected)
		}
	}
}
