//ff:func feature=scan type=test control=sequence topic=flask
//ff:what TestApplyRegisterBlueprintOverrides_NoRegisterCall 테스트
package flask

import "testing"

func TestApplyRegisterBlueprintOverrides_NoRegisterCall(t *testing.T) {
	src := `x = 1
y = foo(2)
`
	fi := flaskFile(t, src)
	prefixes := blueprintPrefix{}
	applyRegisterBlueprintOverrides(fi, prefixes)
	if len(prefixes) != 0 {
		t.Fatalf("expected no overrides, got %v", prefixes)
	}
}
