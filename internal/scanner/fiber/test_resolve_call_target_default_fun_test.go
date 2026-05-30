//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_DefaultFun 테스트
package fiber

import "testing"

func TestResolveCallTarget_DefaultFun(t *testing.T) {

	call := parseCall(t, "(func(){})()")
	if pos := resolveCallTarget(call, newEmptyInfo()); pos.IsValid() {
		t.Fatal("expected NoPos for non-name fun")
	}
}
