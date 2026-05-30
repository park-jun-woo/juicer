//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectViewSetsFromFile_None 테스트
package django

import "testing"

func TestCollectViewSetsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if v := collectViewSetsFromFile(fi); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
