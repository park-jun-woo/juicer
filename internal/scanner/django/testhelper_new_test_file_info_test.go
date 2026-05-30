//ff:func feature=scan type=test control=sequence topic=django
//ff:what newTestFileInfo 테스트 헬퍼
package django

import "testing"

func newTestFileInfo(t *testing.T, src string) fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{relPath: "app/views.py", module: "app.views", src: []byte(src), root: root}
}
