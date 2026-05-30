//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestMergeImportedRouterPrefixes_NoIncludes 테스트
package fastapi

import "testing"

func TestMergeImportedRouterPrefixes_NoIncludes(t *testing.T) {
	src := []byte("x = 1\n")
	root, _ := parsePython(src)
	fi := fileInfo{absPath: "/m.py", src: src, root: root, prefixes: map[string]string{}}
	files := []fileInfo{fi}

	mergeImportedRouterPrefixes("/", files, map[string]map[string]string{})
	if len(files[0].prefixes) != 0 {
		t.Fatalf("expected no change, got %v", files[0].prefixes)
	}
}
