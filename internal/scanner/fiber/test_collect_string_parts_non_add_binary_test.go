//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_NonAddBinary 테스트
package fiber

import "testing"

func TestCollectStringParts_NonAddBinary(t *testing.T) {

	if got := collectFor(t, "a - b"); len(got) != 0 {
		t.Fatalf("expected empty for non-ADD binary, got %v", got)
	}
}
