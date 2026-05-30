//ff:func feature=scan type=test control=sequence
//ff:what TestExtractPathString_NonString 테스트
package fiber

import "testing"

func TestExtractPathString_NonString(t *testing.T) {
	_, ok := extractPathFor(t, "42")
	if ok {
		t.Fatal("int literal should be false")
	}
}
