//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectAPIViewsFromFile_None 테스트
package django

import "testing"

func TestCollectAPIViewsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "def f():\n    pass\n")
	if v := collectAPIViewsFromFile(fi); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
