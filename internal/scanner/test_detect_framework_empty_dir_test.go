//ff:func feature=scan type=test control=sequence
//ff:what TestDetectFramework_EmptyDir 테스트
package scanner

import "testing"

func TestDetectFramework_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	if got := DetectFramework(dir); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
