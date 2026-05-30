//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_NotAdd 테스트
package fiber

import "testing"

func TestExtractBinaryPath_NotAdd(t *testing.T) {
	_, ok := extractBinaryPath(binExpr(t, "a - b"))
	if ok {
		t.Fatal("non-ADD should return false")
	}
}
