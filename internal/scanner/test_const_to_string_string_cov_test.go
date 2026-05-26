//ff:func feature=scan type=test control=sequence
//ff:what TestConstToString_StringCov 테스트
package scanner

import (
	"go/constant"
	"testing"
)

func TestConstToString_StringCov(t *testing.T) {
	v := constant.MakeString("hello")
	got := ConstToString(v)
	if got != `"hello"` {
		t.Fatalf("expected quoted hello, got %s", got)
	}
}
