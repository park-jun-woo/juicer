//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberRouterTypeInfo_ImportedNamed 테스트
package fiber

import "testing"

func TestIsFiberRouterTypeInfo_ImportedNamed(t *testing.T) {

	src := `package m
import "bytes"
var B = bytes.Buffer{}
`
	if isFiberRouterTypeInfo(typeOfVar(t, src, "B")) {
		t.Error("bytes.Buffer should be false")
	}
}
