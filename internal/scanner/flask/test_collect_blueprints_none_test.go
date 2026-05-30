//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestCollectBlueprints_None 테스트
package flask

import "testing"

func TestCollectBlueprints_None(t *testing.T) {
	root, _ := parsePython([]byte("x = 1\n"))
	if bps := collectBlueprints(root, []byte("x = 1\n")); len(bps) != 0 {
		t.Fatalf("expected none, got %d", len(bps))
	}
}
