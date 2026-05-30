//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractParamNameAndType_Round5 테스트
package spring

import "testing"

func TestExtractParamNameAndType_Round5(t *testing.T) {
	param, src := sParam(t, `class C { void m(String userName) {} }`)
	if got := extractParamName(param, src); got != "userName" {
		t.Errorf("name: %q", got)
	}
	if got := extractParamType(param, src); got != "String" {
		t.Errorf("type: %q", got)
	}
}
