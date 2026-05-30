//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractRouterRegistrationsFromFile — router.register() 수집 분기를 검증
package django

import "testing"

func TestExtractRouterRegistrationsFromFile(t *testing.T) {
	src := `
router.register('users', UserViewSet)
foo()
`
	fi := newTestFileInfo(t, src)
	regs := extractRouterRegistrationsFromFile(fi)
	if len(regs) != 1 {
		t.Fatalf("expected 1 registration, got %d", len(regs))
	}
	if regs[0].prefix != "users" {
		t.Errorf("prefix = %q, want users", regs[0].prefix)
	}
}

func TestExtractRouterRegistrationsFromFile_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if r := extractRouterRegistrationsFromFile(fi); len(r) != 0 {
		t.Fatalf("expected none, got %d", len(r))
	}
}
