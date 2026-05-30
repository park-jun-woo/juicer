//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaPropertyToField_Array 테스트
package fastify

import "testing"

func TestJSONSchemaPropertyToField_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array", items: { type: "string" } }`)
	f := jsonSchemaPropertyToField("tags", obj, src)
	if f.Type != "string[]" {
		t.Errorf("expected string[], got %q", f.Type)
	}
}
