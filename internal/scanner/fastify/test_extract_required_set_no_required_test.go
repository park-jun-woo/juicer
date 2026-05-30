//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRequiredSet_NoRequired 테스트
package fastify

import "testing"

func TestExtractRequiredSet_NoRequired(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object" }`)
	if set := extractRequiredSet(obj, src); set != nil {
		t.Fatalf("expected nil, got %v", set)
	}
}
