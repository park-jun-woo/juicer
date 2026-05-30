//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestResolveHTTPMethods_RouteWithMethods 테스트
package flask

import "testing"

func TestResolveHTTPMethods_RouteWithMethods(t *testing.T) {
	args, src := argListOf(t, `route('/x', methods=['POST', 'PUT'])`+"\n")
	got := resolveHTTPMethods("route", args, src)
	if len(got) != 2 || got[0] != "POST" {
		t.Fatalf("route methods: %v", got)
	}
}
