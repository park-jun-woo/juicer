//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractInterfaces_None 테스트
package spring

import "testing"

func TestExtractInterfaces_None(t *testing.T) {
	root, src := parseS(t, `class C {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractInterfaces(cls, src); got != nil {
		t.Fatalf("got %v", got)
	}
}
