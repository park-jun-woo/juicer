//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractClassPrefix 테스트
package spring

import "testing"

func TestExtractClassPrefix(t *testing.T) {
	root, src := parseS(t, `@RequestMapping("/api") class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractClassPrefix(cls, src); got != "/api" {
		t.Fatalf("got %q", got)
	}
}
