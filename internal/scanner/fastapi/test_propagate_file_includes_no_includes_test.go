//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestPropagateFileIncludes_NoIncludes 테스트
package fastapi

import "testing"

func TestPropagateFileIncludes_NoIncludes(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	fi := &fileInfo{absPath: "/m.py", src: src, root: root, prefixes: map[string]string{}}
	if propagateFileIncludes("/", fi, map[string]*fileInfo{}, map[string]map[string]string{}) {
		t.Fatal("expected false when no includes")
	}
}
