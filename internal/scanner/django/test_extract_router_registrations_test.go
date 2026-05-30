//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractRouterRegistrations 테스트
package django

import "testing"

func TestExtractRouterRegistrations(t *testing.T) {
	fi := newTestFileInfo(t, "router.register('users', UserViewSet)\n")
	regs := extractRouterRegistrations([]fileInfo{fi})
	if len(regs) != 1 {
		t.Fatalf("expected 1 registration, got %d", len(regs))
	}
}
