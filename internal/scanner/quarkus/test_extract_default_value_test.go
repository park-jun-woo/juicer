//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractDefaultValue 테스트
package quarkus

import "testing"

func TestExtractDefaultValue(t *testing.T) {
	root, src := parseQ(t, `class R { public void m(@DefaultValue("10") @QueryParam("limit") int limit) {} }`)
	params := findAllByType(root, "formal_parameter")
	if got := extractDefaultValue(params[0], src); got != "10" {
		t.Fatalf("got %q", got)
	}
}
