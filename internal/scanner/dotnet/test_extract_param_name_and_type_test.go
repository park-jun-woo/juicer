//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractParamNameAndType 테스트
package dotnet

import "testing"

func TestExtractParamNameAndType(t *testing.T) {
	root, src := parseCS(t, `class C { void m(int id) {} }`)
	params := findAllByType(root, "parameter")
	if len(params) == 0 {
		t.Fatal("no param")
	}
	if got := extractParamName(params[0], src); got != "id" {
		t.Fatalf("name: %q", got)
	}
	if got := extractParamType(params[0], src); got != "int" {
		t.Fatalf("type: %q", got)
	}
}
