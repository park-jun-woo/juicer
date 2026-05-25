//ff:func feature=scan type=extract control=sequence
//ff:what TestConstToString_String 테스트
package scanner

import (
	"go/constant"
	"testing"
)

func TestConstToString_String(t *testing.T) {
	v := constant.MakeString("hello")
	got := constToString(v)
	if got == "" {
		t.Fatal("expected non-empty")
	}
}
