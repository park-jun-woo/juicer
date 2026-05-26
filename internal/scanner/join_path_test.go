//ff:func feature=scan type=test control=sequence
//ff:what TestJoinPath_Both 테스트
package scanner

import "testing"

func TestJoinPath_Both(t *testing.T) {
	got := joinPath("/api", "/v1")
	if got != "/api/v1" {
		t.Fatalf("expected /api/v1, got %s", got)
	}
}

