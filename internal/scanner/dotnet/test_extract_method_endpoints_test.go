//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestExtractMethodEndpoints 테스트
package dotnet

import "testing"

func TestExtractMethodEndpoints(t *testing.T) {
	fi := csFileInfo(t, sampleCtrl)
	cls := findAllByType(fi.root, "class_declaration")[0]
	eps := extractMethodEndpoints(cls, fi)
	if len(eps) != 2 {
		t.Fatalf("expected 2, got %d", len(eps))
	}
}
