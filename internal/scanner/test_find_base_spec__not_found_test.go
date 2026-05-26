//ff:func feature=scan type=extract control=sequence
//ff:what TestFindBaseSpec_NotFound 테스트
package scanner

import "testing"

func TestFindBaseSpec_NotFound(t *testing.T) {
	dir := t.TempDir()
	got := FindBaseSpec(dir)
	if got != "" {
		t.Fatalf("expected empty string, got %s", got)
	}
}
