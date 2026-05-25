//ff:func feature=scan type=extract control=sequence
//ff:what TestIsRequired_True 테스트
package scanner

import "testing"

func TestIsRequired_True(t *testing.T) {
	if !isRequired(Field{Validate: "required"}) {
		t.Fatal("expected true")
	}
}
