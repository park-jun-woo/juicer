//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRequiredSet_Array 테스트
package fastify

import "testing"

func TestExtractRequiredSet_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object", required: ["name", "email"] }`)
	set := extractRequiredSet(obj, src)
	if !set["name"] || !set["email"] {
		t.Fatalf("expected name+email, got %v", set)
	}
	if len(set) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(set))
	}
}
