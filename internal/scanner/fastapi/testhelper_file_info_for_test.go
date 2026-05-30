//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what fileInfoFor 테스트 헬퍼
package fastapi

import "testing"

func fileInfoFor(t *testing.T, src string) fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{src: []byte(src), root: root}
}
