//ff:func feature=scan type=test control=sequence
//ff:what TestJoinPath_EmptyACov 테스트
package scanner

import "testing"

func TestJoinPath_EmptyACov(t *testing.T) {
	got := JoinPath("", "/v1")
	if got != "/v1" {
		t.Fatalf("expected /v1, got %s", got)
	}
}
