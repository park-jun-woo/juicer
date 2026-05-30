//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestCollectFuncViewsFromFile_None 테스트
package django

import "testing"

func TestCollectFuncViewsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if v := collectFuncViewsFromFile(fi); len(v) != 0 {
		t.Fatalf("expected none, got %d", len(v))
	}
}
