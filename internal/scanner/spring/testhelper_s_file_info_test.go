//ff:func feature=scan type=test control=sequence topic=spring
//ff:what sFileInfo 테스트 헬퍼
package spring

import "testing"

func sFileInfo(t *testing.T, javaSrc string) *fileInfo {
	t.Helper()
	b := []byte(javaSrc)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return &fileInfo{
		absPath:     "/abs/C.java",
		relPath:     "C.java",
		projectRoot: "/abs",
		src:         b,
		root:        root,
		imports:     map[string]string{},
	}
}
