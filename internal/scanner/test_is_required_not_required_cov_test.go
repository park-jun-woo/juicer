//ff:func feature=scan type=test control=sequence
//ff:what TestIsRequired_NotRequiredCov 테스트
package scanner

import "testing"

func TestIsRequired_NotRequiredCov(t *testing.T) {
	if isRequired(Field{Validate: "email,min=3"}) {
		t.Fatal("expected false")
	}
}
