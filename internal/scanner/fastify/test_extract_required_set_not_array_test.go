//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractRequiredSet_NotArray 테스트
package fastify

import "testing"

func TestExtractRequiredSet_NotArray(t *testing.T) {
	obj, src := firstObject(t, `{ required: "name" }`)
	if set := extractRequiredSet(obj, src); set != nil {
		t.Fatalf("expected nil when required not array, got %v", set)
	}
}
