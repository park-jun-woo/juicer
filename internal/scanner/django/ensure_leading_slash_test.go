//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what ensureLeadingSlash — 선행 슬래시 보장 분기를 검증
package django

import "testing"

func TestEnsureLeadingSlash(t *testing.T) {
	cases := map[string]string{
		"":         "/",
		"api/x":    "/api/x",
		"/already": "/already",
	}
	for in, want := range cases {
		if got := ensureLeadingSlash(in); got != want {
			t.Errorf("ensureLeadingSlash(%q) = %q, want %q", in, got, want)
		}
	}
}
