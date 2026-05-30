//ff:func feature=scan type=test control=sequence topic=zod
//ff:what TestParseChain 테스트
package zod

import "testing"

func TestParseChain(t *testing.T) {
	root, src := parseTS(t, `const x = z.string().email().min(3);`)
	calls := findAllByType(root, "call_expression")

	f := ParseChain(calls[0], src)
	if f.Type != "string" {
		t.Fatalf("type: %+v", f)
	}
	if f.MinLength == nil || *f.MinLength != 3 {
		t.Fatalf("min: %v", f.MinLength)
	}
}
