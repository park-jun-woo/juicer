//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestFindStaticFinalField 테스트
package spring

import "testing"

func TestFindStaticFinalField(t *testing.T) {
	root, src := parseS(t, `class C { public static final String PREFIX = "api"; }`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := findStaticFinalField(cls, src, "PREFIX"); got != "api" {
		t.Fatalf("got %q", got)
	}
	if got := findStaticFinalField(cls, src, "MISSING"); got != "" {
		t.Fatalf("missing: %q", got)
	}
}
