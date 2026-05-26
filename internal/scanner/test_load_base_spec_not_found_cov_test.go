//ff:func feature=scan type=test control=sequence
//ff:what TestLoadBaseSpec_NotFoundCov 테스트
package scanner

import "testing"

func TestLoadBaseSpec_NotFoundCov(t *testing.T) {
	_, err := LoadBaseSpec("/nonexistent/file.yaml")
	if err == nil {
		t.Fatal("expected error")
	}
}
