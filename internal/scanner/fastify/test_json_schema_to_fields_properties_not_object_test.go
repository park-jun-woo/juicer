//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToFields_PropertiesNotObject 테스트
package fastify

import "testing"

func TestJSONSchemaToFields_PropertiesNotObject(t *testing.T) {
	obj, src := firstObject(t, `{ type: "object", properties: "x" }`)
	if got := jsonSchemaToFields(obj, src); got != nil {
		t.Fatalf("expected nil when properties not object, got %v", got)
	}
}
