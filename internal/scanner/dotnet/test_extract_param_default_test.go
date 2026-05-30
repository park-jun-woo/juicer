//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractParamDefault 테스트
package dotnet

import "testing"

func TestExtractParamDefault(t *testing.T) {
	root, src := parseCS(t, `class C { void m(int limit = 10) {} }`)
	params := findAllByType(root, "parameter")
	if got := extractParamDefault(params[0], src); got != "10" {
		t.Fatalf("got %q", got)
	}
}
