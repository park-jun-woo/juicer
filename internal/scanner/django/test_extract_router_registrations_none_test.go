//ff:func feature=scan type=test control=sequence topic=django
//ff:what TestExtractRouterRegistrations_None 테스트
package django

import "testing"

func TestExtractRouterRegistrations_None(t *testing.T) {
	fi := newTestFileInfo(t, "x = 1\n")
	if r := extractRouterRegistrations([]fileInfo{fi}); len(r) != 0 {
		t.Fatalf("expected none, got %d", len(r))
	}
}
