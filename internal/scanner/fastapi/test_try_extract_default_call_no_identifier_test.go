//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what TestTryExtractDefault_CallNoIdentifier 테스트
package fastapi

import "testing"

func TestTryExtractDefault_CallNoIdentifier(t *testing.T) {

	src := []byte("def f(x = mod.factory()): pass\n")
	root, _ := parsePython(src)
	c := firstOfType(root, "call")
	if c == nil {
		t.Skip("no call node")
	}
	v, call, none := tryExtractDefault(c, src)
	if v == "" || none {
		t.Fatalf("call val should be non-empty: v=%q none=%v", v, none)
	}
	_ = call
}
