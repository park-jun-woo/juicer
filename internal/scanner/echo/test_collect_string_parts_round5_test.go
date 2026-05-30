//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestCollectStringParts_Round5 테스트
package echo

import "testing"

func TestCollectStringParts_Round5(t *testing.T) {
	var parts []string
	collectStringParts(nil, parseExpr(t, `"a" + "b" + "c"`), &parts)
	if len(parts) != 3 || parts[0] != "a" || parts[2] != "c" {
		t.Fatalf("parts: %v", parts)
	}
}
