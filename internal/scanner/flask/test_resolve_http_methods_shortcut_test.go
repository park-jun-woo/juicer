//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestResolveHTTPMethods_Shortcut 테스트
package flask

import "testing"

func TestResolveHTTPMethods_Shortcut(t *testing.T) {
	args, src := argListOf(t, `get('/x')`+"\n")
	got := resolveHTTPMethods("get", args, src)
	if len(got) != 1 || got[0] != "GET" {
		t.Fatalf("get shortcut: %v", got)
	}
}
