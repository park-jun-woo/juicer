//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestParseRegisterCall_Round5 테스트
package django

import "testing"

func TestParseRegisterCall_Round5(t *testing.T) {
	fi := newTestFileInfo(t, "router.register('users', UserViewSet)\n")
	call := djFirst(t, fi.root, "call")
	reg := parseRegisterCall(call, fi)
	if reg == nil {
		t.Fatal("expected registration")
	}
}
