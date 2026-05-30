//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateSingleInclude_SrcFINil 테스트
package fastapi

import "testing"

func TestPropagateSingleInclude_SrcFINil(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{})
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src.py"}

	if propagateSingleInclude(fi, inc, importMap, map[string]*fileInfo{}, map[string]map[string]string{}) {
		t.Fatal("expected false when srcFI nil")
	}
}
