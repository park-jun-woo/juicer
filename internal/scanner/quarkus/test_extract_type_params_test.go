//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractTypeParams 테스트
package quarkus

import "testing"

func TestExtractTypeParams(t *testing.T) {
	root, src := parseQ(t, `class R<T, U> {}`)
	cls := findAllByType(root, "class_declaration")[0]
	got := extractTypeParams(cls, src)
	if len(got) != 2 || got[0] != "T" {
		t.Fatalf("got %v", got)
	}
}
