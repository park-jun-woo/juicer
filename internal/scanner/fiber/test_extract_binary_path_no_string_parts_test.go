//ff:func feature=scan type=test control=sequence
//ff:what TestExtractBinaryPath_NoStringParts 테스트
package fiber

import "testing"

func TestExtractBinaryPath_NoStringParts(t *testing.T) {

	_, ok := extractBinaryPath(binExpr(t, "a + b"))
	if ok {
		t.Fatal("ident concat should return false")
	}
}
