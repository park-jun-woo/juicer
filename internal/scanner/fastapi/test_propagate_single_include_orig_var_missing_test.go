//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateSingleInclude_OrigVarMissing 테스트
package fastapi

import "testing"

func TestPropagateSingleInclude_OrigVarMissing(t *testing.T) {
	fi := newFileInfoSimple(t, "x=1\n", map[string]string{})
	srcFI := newFileInfoSimple(t, "x=1\n", map[string]string{})
	srcFI.absPath = "/src.py"
	inc := includeCall{parentVar: "app", childVar: "users"}
	importMap := map[string]string{"users": "/src.py"}
	fileByPath := map[string]*fileInfo{"/src.py": srcFI}
	if propagateSingleInclude(fi, inc, importMap, fileByPath, map[string]map[string]string{}) {
		t.Fatal("expected false when origVar missing")
	}
}
