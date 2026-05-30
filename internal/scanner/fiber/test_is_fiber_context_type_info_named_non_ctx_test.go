//ff:func feature=scan type=test control=sequence
//ff:what TestIsFiberContextTypeInfo_NamedNonCtx 테스트
package fiber

import "testing"

func TestIsFiberContextTypeInfo_NamedNonCtx(t *testing.T) {

	src := `package m
import "bytes"
var B = &bytes.Buffer{}
`
	if isFiberContextTypeInfo(typeOfVar(t, src, "B")) {
		t.Error("*bytes.Buffer should be false")
	}
}
