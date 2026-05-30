//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractClassPrefix_None 테스트
package spring

import "testing"

func TestExtractClassPrefix_None(t *testing.T) {
	root, src := parseS(t, `class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractClassPrefix(cls, src); got != "" {
		t.Fatalf("got %q", got)
	}
}
