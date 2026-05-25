//ff:func feature=scan type=extract control=sequence
//ff:what TestIsRequired_Empty 테스트
package scanner

import "testing"

func TestIsRequired_Empty(t *testing.T) {
	if isRequired(Field{}) {
		t.Fatal("expected false for empty validate")
	}
}
