//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToFields_NoProperties 테스트
package fastify

import "testing"

func TestJSONSchemaToFields_NoProperties(t *testing.T) {

	obj, src := firstObject(t, `{ type: "object" }`)
	if got := jsonSchemaToFields(obj, src); got != nil {
		t.Fatalf("expected nil when no properties, got %v", got)
	}
}
