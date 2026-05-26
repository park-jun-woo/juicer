//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what filterNone 테스트
package fastapi

import "testing"

func TestFilterNone(t *testing.T) {
	got := filterNone([]string{"str", "None", " None ", "int"})
	if len(got) != 2 || got[0] != "str" || got[1] != "int" {
		t.Fatalf("unexpected: %v", got)
	}
	got2 := filterNone(nil)
	if len(got2) != 0 {
		t.Fatalf("expected empty, got %v", got2)
	}
}
