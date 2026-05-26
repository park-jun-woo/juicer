//ff:func feature=scan type=test control=sequence
//ff:what TestFindBaseSpec_NotFoundCov 테스트
package scanner

import "testing"

func TestFindBaseSpec_NotFoundCov(t *testing.T) {
	dir := t.TempDir()
	result := FindBaseSpec(dir)
	if result != "" {
		t.Fatalf("expected empty, got %s", result)
	}
}
