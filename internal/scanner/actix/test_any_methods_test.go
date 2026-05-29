//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what TestAnyMethods — method 미지정 리소스가 5개 HTTP 메서드로 전개되는지 검증
package actix

import "testing"

func TestAnyMethods(t *testing.T) {
	got := anyMethods()
	want := []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	if len(got) != len(want) {
		t.Fatalf("len: got %d want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d]: got %q want %q", i, got[i], want[i])
		}
	}
}
