//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what 테스트 헬퍼: 응답 집합이 기대값과 일치하는지 검증한다
package express

import (
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func assertResponseSet(t *testing.T, got []scanner.Response, want []scanner.Response) {
	t.Helper()
	gotSet := map[string]bool{}
	for _, r := range got {
		gotSet[r.Status+"/"+r.Kind] = true
	}
	for _, w := range want {
		key := w.Status + "/" + w.Kind
		if !gotSet[key] {
			t.Errorf("missing response %s; got %v", key, got)
		}
	}
	if len(got) != len(want) {
		t.Errorf("response count: want %d, got %d (%v)", len(want), len(got), got)
	}
}
