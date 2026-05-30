//ff:func feature=scan type=test control=sequence
//ff:what TestResolveCallTarget_Unresolved 테스트
package fiber

import "testing"

func TestResolveCallTarget_Unresolved(t *testing.T) {

	call := parseCall(t, "unknown()")
	if pos := resolveCallTarget(call, newEmptyInfo()); pos.IsValid() {
		t.Fatal("expected NoPos for unresolved ident")
	}
}
