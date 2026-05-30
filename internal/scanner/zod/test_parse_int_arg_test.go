//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestParseIntArg 테스트
package zod

import "testing"

func TestParseIntArg(t *testing.T) {
	if p := parseIntArg("5"); p == nil || *p != 5 {
		t.Fatal("5")
	}
	if parseIntArg("x") != nil {
		t.Fatal("invalid")
	}
}
