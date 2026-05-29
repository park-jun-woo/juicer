//ff:func feature=scan type=test control=iteration dimension=1 topic=dotnet
//ff:what TestStripRouteConstraints -- 라우트 제약/기본값/옵셔널 토큰 정규화 테스트
package dotnet

import "testing"

func TestStripRouteConstraints(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"api/v{version:apiVersion}/auth", "api/v{version}/auth"},
		{"items/{id:int}", "items/{id}"},
		{"items/{id:int:min(1)}", "items/{id}"},
		{"users/{name?}", "users/{name}"},
		{"posts/{slug=home}", "posts/{slug}"},
		{"files/{*slug}", "files/{slug}"},
		{"files/{**slug}", "files/{slug}"},
		{"plain/{id}", "plain/{id}"},
		{"no/tokens/here", "no/tokens/here"},
	}
	for _, c := range cases {
		if got := stripRouteConstraints(c.in); got != c.want {
			t.Errorf("stripRouteConstraints(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
