//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectViewSets_None 테스트
package django

import "testing"

func TestCollectViewSets_None(t *testing.T) {
	fi := newTestFileInfo(t, "class Plain:\n    pass\n")
	if v := collectViewSets([]fileInfo{fi}); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
