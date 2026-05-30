//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateSingleInclude_SrcFileMissing 테스트
package fastapi

import "testing"

func TestPropagateSingleInclude_SrcFileMissing(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{})
	inc := includeCall{parentVar: "app", childVar: "users"}
	if propagateSingleInclude(fi, inc, map[string]string{}, map[string]*fileInfo{}, map[string]map[string]string{}) {
		t.Fatal("expected false when srcFile missing")
	}
}
