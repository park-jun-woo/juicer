//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what qFileInfo 테스트 헬퍼
package quarkus

import "testing"

func qFileInfo(t *testing.T, javaSrc string) *fileInfo {
	t.Helper()
	b := []byte(javaSrc)
	root, err := parseJava(b)
	if err != nil {
		t.Fatal(err)
	}
	return &fileInfo{
		absPath:     "/abs/R.java",
		relPath:     "R.java",
		projectRoot: "/abs",
		src:         b,
		root:        root,
		imports:     map[string]string{},
	}
}
