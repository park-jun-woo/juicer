//ff:func feature=scan type=test control=iteration dimension=1 topic=express
//ff:what assertHas — 기대 엔드포인트 키들이 모두 존재하는지 검증한다
package express

import "testing"

func assertHas(t *testing.T, got map[string]bool, want ...string) {
	t.Helper()
	for _, w := range want {
		if !got[w] {
			t.Errorf("missing endpoint %q; got %v", w, got)
		}
	}
}
