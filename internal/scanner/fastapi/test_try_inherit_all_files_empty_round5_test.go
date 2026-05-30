//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryInheritAllFiles_Empty_Round5 테스트
package fastapi

import "testing"

func TestTryInheritAllFiles_Empty_Round5(t *testing.T) {
	if tryInheritAllFiles(nil, map[string]*fileInfo{}) {
		t.Fatal("empty files should return false")
	}
}
