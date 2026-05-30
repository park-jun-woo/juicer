//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestResolveHTTPMethods_RouteDefault 테스트
package flask

import "testing"

func TestResolveHTTPMethods_RouteDefault(t *testing.T) {
	args, src := argListOf(t, `route('/x')`+"\n")
	got := resolveHTTPMethods("route", args, src)
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("route default: %v", got)
	}
}
