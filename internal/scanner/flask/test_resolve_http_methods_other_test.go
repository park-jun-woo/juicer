//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestResolveHTTPMethods_Other 테스트
package flask

import "testing"

func TestResolveHTTPMethods_Other(t *testing.T) {
	args, src := argListOf(t, `before_request('x')`+"\n")
	if got := resolveHTTPMethods("before_request", args, src); got != nil {
		t.Fatalf("non-route method should be nil, got %v", got)
	}
}
