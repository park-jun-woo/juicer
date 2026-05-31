//ff:func feature=scan type=test control=sequence topic=django
//ff:what mkFile — relPath/module 지정 fileInfo 생성 테스트 헬퍼
package django

import "testing"

func mkFile(t *testing.T, relPath, module, src string) fileInfo {
	t.Helper()
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	return fileInfo{relPath: relPath, module: module, src: []byte(src), root: root}
}
