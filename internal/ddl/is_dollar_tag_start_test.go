//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what isDollarTagStart의 letter/underscore 허용 및 숫자/기타 거부 테스트
package ddl

import "testing"

func TestIsDollarTagStart(t *testing.T) {
	starts := []rune{'a', 'z', 'A', 'Z', '_'}
	for _, ch := range starts {
		if !isDollarTagStart(ch) {
			t.Errorf("expected %q to be a valid start char", ch)
		}
	}

	notStarts := []rune{'0', '9', '$', '-', ' '}
	for _, ch := range notStarts {
		if isDollarTagStart(ch) {
			t.Errorf("expected %q to be rejected as start char", ch)
		}
	}
}
