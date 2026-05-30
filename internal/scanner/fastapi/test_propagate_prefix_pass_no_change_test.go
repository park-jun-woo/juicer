//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagatePrefixPass_NoChange 테스트
package fastapi

import "testing"

func TestPropagatePrefixPass_NoChange(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	files := []fileInfo{
		{absPath: "/a.py", src: src, root: root, prefixes: map[string]string{}},
		{absPath: "/b.py", src: src, root: root, prefixes: map[string]string{}},
	}
	fileByPath := map[string]*fileInfo{}
	if propagatePrefixPass("/", files, fileByPath, map[string]map[string]string{}) {
		t.Fatal("expected false when nothing to propagate")
	}
}
