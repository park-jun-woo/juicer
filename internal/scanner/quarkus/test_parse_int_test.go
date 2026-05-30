//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestParseInt 테스트
package quarkus

import "testing"

func TestParseInt(t *testing.T) {
	if parseInt("123") != 123 {
		t.Fatal("123")
	}
	if parseInt("a1b2") != 12 {
		t.Fatal("a1b2")
	}
	if parseInt("") != 0 {
		t.Fatal("empty")
	}
}
