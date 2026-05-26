//ff:func feature=scan type=test control=sequence
//ff:what TestIsRequired_EmptyCov 테스트
package scanner

import "testing"

func TestIsRequired_EmptyCov(t *testing.T) {
	if isRequired(Field{}) {
		t.Fatal("expected false for empty validate")
	}
}
