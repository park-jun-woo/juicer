//ff:func feature=hurl type=parse control=sequence
//ff:what TestFirstTODO_Empty 테스트
package hurls

import "testing"

func TestFirstTODO_Empty(t *testing.T) {
	sess := &Session{}
	if got := firstTODO(sess); got != -1 {
		t.Fatalf("expected -1, got %d", got)
	}
}
