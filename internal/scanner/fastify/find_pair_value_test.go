//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what findPairValue 테스트
package fastify

import "testing"

func TestFindPairValue(t *testing.T) {
	obj, src := firstObject(t, `{ a: "x", b: 2 }`)
	val := findPairValue(obj, src, "a")
	if val == nil || val.Type() != "string" {
		t.Fatalf("expected string value for a, got %v", val)
	}
	// missing key -> nil
	if got := findPairValue(obj, src, "missing"); got != nil {
		t.Fatalf("expected nil for missing key, got %s", got.Type())
	}
}
