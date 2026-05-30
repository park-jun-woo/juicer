//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestUnquoteCSharp 테스트
package dotnet

import "testing"

func TestUnquoteCSharp(t *testing.T) {
	if unquoteCSharp(`"x"`) != "x" || unquoteCSharp("y") != "y" {
		t.Fatal("unquote")
	}
}
