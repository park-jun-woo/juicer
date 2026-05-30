//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestParseInt 테스트
package spring

import "testing"

func TestParseInt(t *testing.T) {
	if parseInt("204") != 204 || parseInt("x9y") != 9 || parseInt("") != 0 {
		t.Fatal("parseInt")
	}
}
