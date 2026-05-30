//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what mkParsedFile 테스트 헬퍼
package fastapi

import "testing"

func mkParsedFile(t *testing.T, src string) fileInfo {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{src: b, root: root}
}
