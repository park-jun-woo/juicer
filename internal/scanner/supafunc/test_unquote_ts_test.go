//ff:func feature=scan type=test control=sequence topic=supafunc
//ff:what TestUnquoteTS 테스트
package supafunc

import "testing"

func TestUnquoteTS(t *testing.T) {
	if unquoteTS(`"x"`) != "x" || unquoteTS("'y'") != "y" || unquoteTS("`z`") != "z" || unquoteTS("a") != "a" {
		t.Fatal("unquote")
	}
}
