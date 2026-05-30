//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestUnquoteJava 테스트
package quarkus

import "testing"

func TestUnquoteJava(t *testing.T) {
	if unquoteJava(`"hi"`) != "hi" {
		t.Fatal("quoted")
	}
	if unquoteJava("x") != "x" {
		t.Fatal("short")
	}
	if unquoteJava("noq") != "noq" {
		t.Fatal("unquoted")
	}
}
