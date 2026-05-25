//ff:func feature=hurl type=parse control=sequence
//ff:what TestFindTestFile_NoDir 테스트
package hurls

import "testing"

func TestFindTestFile_NoDir(t *testing.T) {
	got := findTestFile("/nonexistent/dir", "GET", "/path")
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
