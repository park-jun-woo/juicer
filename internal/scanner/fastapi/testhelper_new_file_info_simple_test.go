//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what newFileInfoSimple 테스트 헬퍼
package fastapi

import "testing"

func newFileInfoSimple(t *testing.T, src string, prefixes map[string]string) *fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return &fileInfo{absPath: "/main.py", src: []byte(src), root: root, prefixes: prefixes}
}
