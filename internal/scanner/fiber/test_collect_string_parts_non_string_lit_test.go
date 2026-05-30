//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_NonStringLit 테스트
package fiber

import "testing"

func TestCollectStringParts_NonStringLit(t *testing.T) {

	if got := collectFor(t, "42"); len(got) != 0 {
		t.Fatalf("expected empty for int lit, got %v", got)
	}
}
