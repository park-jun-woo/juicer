//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestUnquoteJava 테스트
package spring

import "testing"

func TestUnquoteJava(t *testing.T) {
	if unquoteJava(`"x"`) != "x" || unquoteJava("y") != "y" {
		t.Fatal("unquote")
	}
}
