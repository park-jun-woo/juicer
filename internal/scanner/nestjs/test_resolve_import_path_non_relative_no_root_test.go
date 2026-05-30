//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_NonRelativeNoRoot 테스트
package nestjs

import "testing"

func TestResolveImportPath_NonRelativeNoRoot(t *testing.T) {
	if got := resolveImportPath("/x", "@nestjs/common"); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
