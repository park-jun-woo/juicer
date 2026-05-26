//ff:func feature=scan type=test control=sequence
//ff:what TestConstToString_Int 테스트
package scanner

import (
	"go/constant"
	"testing"
)

func TestConstToString_Int(t *testing.T) {
	v := constant.MakeInt64(200)
	got := ConstToString(v)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}

	// non-int constant (string)
	sv := constant.MakeString("hello")
	got = ConstToString(sv)
	if got != `"hello"` {
		t.Fatalf("expected \"hello\", got %s", got)
	}
}

