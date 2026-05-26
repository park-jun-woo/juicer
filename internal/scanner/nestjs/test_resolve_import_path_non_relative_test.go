//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestResolveImportPath_NonRelative 테스트
package nestjs

import "testing"

func TestResolveImportPath_NonRelative(t *testing.T) {
	result := resolveImportPath("/tmp", "@nestjs/common")
	if result != "" {
		t.Fatal("expected empty for non-relative import")
	}
}
