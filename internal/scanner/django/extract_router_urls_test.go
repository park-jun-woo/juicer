//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractRouterRegistrations — 여러 파일의 router.register() 수집을 검증
package django

import "testing"

func TestExtractRouterRegistrations(t *testing.T) {
	fi := newTestFileInfo(t, "router.register('users', UserViewSet)\n")
	regs := extractRouterRegistrations([]fileInfo{fi})
	if len(regs) != 1 {
		t.Fatalf("expected 1 registration, got %d", len(regs))
	}
}

func TestExtractRouterRegistrations_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if r := extractRouterRegistrations([]fileInfo{fi}); len(r) != 0 {
		t.Fatalf("expected none, got %d", len(r))
	}
}
