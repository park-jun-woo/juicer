//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what isDollarTagCont의 letter/underscore/digit 허용 및 기타 문자 거부 테스트
package ddl

import "testing"

func TestIsDollarTagCont(t *testing.T) {
	cont := []rune{'a', 'Z', '_', '0', '9'}
	for _, ch := range cont {
		if !isDollarTagCont(ch) {
			t.Errorf("expected %q to be a valid continuation char", ch)
		}
	}

	notCont := []rune{'$', '-', ' ', '.'}
	for _, ch := range notCont {
		if isDollarTagCont(ch) {
			t.Errorf("expected %q to be rejected", ch)
		}
	}
}
