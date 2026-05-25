//ff:func feature=scan type=extract control=sequence
//ff:what TestIsRequired_InComma 테스트
package scanner

import "testing"

func TestIsRequired_InComma(t *testing.T) {
	if !isRequired(Field{Validate: "required,email"}) {
		t.Fatal("expected true")
	}
}
