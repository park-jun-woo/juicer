//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractSuperclassName 테스트
package quarkus

import "testing"

func TestExtractSuperclassName(t *testing.T) {
	root, src := parseQ(t, `class Child extends Base {}`)
	sc := findAllByType(root, "superclass")
	if len(sc) == 0 {
		t.Skip("no superclass")
	}
	if got := extractSuperclassName(sc[0], src); got != "Base" {
		t.Fatalf("got %q", got)
	}
}
