//ff:func feature=scan type=extract control=sequence
//ff:what TestLoadBaseSpec_NotFound 테스트
package scanner

import "testing"

func TestLoadBaseSpec_NotFound(t *testing.T) {
	_, err := LoadBaseSpec("/nonexistent/path/openapi.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}
