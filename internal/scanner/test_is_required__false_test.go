//ff:func feature=scan type=extract control=sequence
//ff:what TestIsRequired_False 테스트
package scanner

import "testing"

func TestIsRequired_False(t *testing.T) {
	if isRequired(Field{Validate: "email"}) {
		t.Fatal("expected false")
	}
}
