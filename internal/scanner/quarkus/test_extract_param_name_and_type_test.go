//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractParamNameAndType 테스트
package quarkus

import "testing"

func TestExtractParamNameAndType(t *testing.T) {
	root, src := parseQ(t, `class R { public void m(String name) {} }`)
	params := findAllByType(root, "formal_parameter")
	if len(params) == 0 {
		t.Fatal("no params")
	}
	if got := extractParamName(params[0], src); got != "name" {
		t.Fatalf("name: %q", got)
	}
	if got := extractParamType(params[0], src); got != "String" {
		t.Fatalf("type: %q", got)
	}
}
