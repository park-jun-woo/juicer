//ff:func feature=scan type=test control=sequence topic=flask
//ff:what flaskFile 테스트 헬퍼
package flask

import "testing"

func flaskFile(t *testing.T, src string) fileInfo {
	t.Helper()
	b := []byte(src)
	root, err := parsePython(b)
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{src: b, root: root}
}
