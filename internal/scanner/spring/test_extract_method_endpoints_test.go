//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractMethodEndpoints 테스트
package spring

import "testing"

func TestExtractMethodEndpoints(t *testing.T) {
	fi := sFileInfo(t, sampleController)
	cls := findAllByType(fi.root, "class_declaration")[0]
	eps := extractMethodEndpoints(cls, fi)
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}
