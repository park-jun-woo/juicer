//ff:func feature=scan type=extract control=sequence
//ff:what TestLcFirst_Acronym 테스트
package scanner

import "testing"

func TestLcFirst_Acronym(t *testing.T) {
	if lcFirst("ID") != "id" {
		t.Fatal("expected id")
	}
}
