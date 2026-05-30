//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractTypeParams_None 테스트
package quarkus

import "testing"

func TestExtractTypeParams_None(t *testing.T) {
	root, src := parseQ(t, `class R {}`)
	cls := findAllByType(root, "class_declaration")[0]
	if got := extractTypeParams(cls, src); got != nil {
		t.Fatalf("got %v", got)
	}
}
