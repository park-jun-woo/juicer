//ff:func feature=prisma type=test topic=prisma control=sequence
//ff:what defaultArg @default(...) 내부 추출 테스트
package prisma

import "testing"

func TestDefaultArg(t *testing.T) {
	v, ok := defaultArg("@default( now() )")
	if !ok || v != "now()" {
		t.Errorf("got (%q,%v), want (now(),true)", v, ok)
	}
	if _, ok := defaultArg("@id"); ok {
		t.Error("non-default must be false")
	}
	// no closing paren: keeps remainder, trimmed
	v, ok = defaultArg("@default(5")
	if !ok || v != "5" {
		t.Errorf("no close paren: got (%q,%v)", v, ok)
	}
}
