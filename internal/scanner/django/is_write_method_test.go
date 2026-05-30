//ff:func feature=scan type=test control=iteration dimension=1 topic=django
//ff:what isWriteMethod — 쓰기 메서드 판별을 검증
package django

import "testing"

func TestIsWriteMethod(t *testing.T) {
	for _, m := range []string{"POST", "PUT", "PATCH"} {
		if !isWriteMethod(m) {
			t.Errorf("expected %s to be a write method", m)
		}
	}
	for _, m := range []string{"GET", "DELETE", "HEAD", ""} {
		if isWriteMethod(m) {
			t.Errorf("expected %s to not be a write method", m)
		}
	}
}
