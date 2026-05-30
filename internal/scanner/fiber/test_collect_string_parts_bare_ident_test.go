//ff:func feature=scan type=test control=sequence
//ff:what TestCollectStringParts_BareIdent 테스트
package fiber

import "testing"

func TestCollectStringParts_BareIdent(t *testing.T) {
	if got := collectFor(t, "baseURL"); len(got) != 0 {
		t.Fatalf("expected empty for ident, got %v", got)
	}
}
