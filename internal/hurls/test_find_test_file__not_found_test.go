//ff:func feature=hurl type=parse control=sequence
//ff:what TestFindTestFile_NotFound 테스트
package hurls

import "testing"

func TestFindTestFile_NotFound(t *testing.T) {
	dir := t.TempDir()
	got := findTestFile(dir, "GET", "/api/missing")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
