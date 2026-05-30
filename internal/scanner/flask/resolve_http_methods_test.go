//ff:func feature=scan type=test control=sequence topic=flask
//ff:what resolveHTTPMethods 테스트
package flask

import "testing"

func TestResolveHTTPMethods_Shortcut(t *testing.T) {
	args, src := argListOf(t, `get('/x')`+"\n")
	got := resolveHTTPMethods("get", args, src)
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("get shortcut: %v", got)
	}
}

func TestResolveHTTPMethods_RouteWithMethods(t *testing.T) {
	args, src := argListOf(t, `route('/x', methods=['POST', 'PUT'])`+"\n")
	got := resolveHTTPMethods("route", args, src)
	if len(got) != 2 || got[0] != "POST" {
		t.Fatalf("route methods: %v", got)
	}
}

func TestResolveHTTPMethods_RouteDefault(t *testing.T) {
	args, src := argListOf(t, `route('/x')`+"\n")
	got := resolveHTTPMethods("route", args, src)
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("route default: %v", got)
	}
}

func TestResolveHTTPMethods_Other(t *testing.T) {
	args, src := argListOf(t, `before_request('x')`+"\n")
	if got := resolveHTTPMethods("before_request", args, src); got != nil {
		t.Fatalf("non-route method should be nil, got %v", got)
	}
}
