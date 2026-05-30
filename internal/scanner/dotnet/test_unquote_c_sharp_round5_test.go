//ff:func feature=scan type=test control=sequence topic=dotnet
//ff:what TestUnquoteCSharp_Round5 테스트
package dotnet

import "testing"

func TestUnquoteCSharp_Round5(t *testing.T) {
	if got := unquoteCSharp(`"hi"`); got != "hi" {
		t.Errorf("quoted: got %q", got)
	}
	if got := unquoteCSharp("x"); got != "x" {
		t.Errorf("short: got %q", got)
	}
	if got := unquoteCSharp("ab"); got != "ab" {
		t.Errorf("no-quotes: got %q", got)
	}
	if got := unquoteCSharp(""); got != "" {
		t.Errorf("empty: got %q", got)
	}
}
