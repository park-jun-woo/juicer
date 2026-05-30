//ff:func feature=scan type=test control=sequence topic=actix
//ff:what aFi 테스트 헬퍼
package actix

import "testing"

func aFi(t *testing.T, src string) *fileInfo {
	t.Helper()
	root, b := aParse(t, src)
	return &fileInfo{absPath: "/abs/m.rs", relPath: "m.rs", projectRoot: "/abs", src: b, root: root}
}
