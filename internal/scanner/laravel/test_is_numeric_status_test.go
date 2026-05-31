//ff:func feature=scan type=test control=iteration dimension=1 topic=laravel
//ff:what isNumericStatus 숫자 문자열 판별 테스트
package laravel

import "testing"

func TestIsNumericStatus(t *testing.T) {
	for _, s := range []string{"200", "404", "0"} {
		if !isNumericStatus(s) {
			t.Errorf("%q should be numeric", s)
		}
	}
	for _, s := range []string{"", "20a", "abc", "2.0"} {
		if isNumericStatus(s) {
			t.Errorf("%q should not be numeric", s)
		}
	}
}
