//ff:func feature=scan type=extract control=sequence
//ff:what TestJoinPath_EmptyA 테스트
package scanner

import "testing"

func TestJoinPath_EmptyA(t *testing.T) {
	got := JoinPath("", "/v1")
	if got != "/v1" {
		t.Fatalf("expected /v1, got %s", got)
	}
}
