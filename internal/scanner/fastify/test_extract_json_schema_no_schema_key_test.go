//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractJSONSchema_NoSchemaKey 테스트
package fastify

import "testing"

func TestExtractJSONSchema_NoSchemaKey(t *testing.T) {
	obj, src := firstObject(t, `{ config: { x: 1 } }`)
	if si := extractJSONSchema(obj, src); si != nil {
		t.Fatalf("expected nil when no schema key, got %v", si)
	}
}
