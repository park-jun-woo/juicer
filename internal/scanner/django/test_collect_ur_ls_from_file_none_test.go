//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectURLsFromFile_None 테스트
package django

import "testing"

func TestCollectURLsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if e := collectURLsFromFile(fi); len(e) != 0 {
		t.Fatalf("expected none, got %d", len(e))
	}
}
