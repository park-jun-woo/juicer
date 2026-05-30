//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestExtractJSONSchema_SchemaNotObject 테스트
package fastify

import "testing"

func TestExtractJSONSchema_SchemaNotObject(t *testing.T) {

	obj, src := firstObject(t, `{ schema: true }`)
	if si := extractJSONSchema(obj, src); si != nil {
		t.Fatalf("expected nil when schema not object, got %v", si)
	}
}
