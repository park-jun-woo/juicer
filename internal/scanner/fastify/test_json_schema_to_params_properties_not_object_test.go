//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaToParams_PropertiesNotObject 테스트
package fastify

import "testing"

func TestJSONSchemaToParams_PropertiesNotObject(t *testing.T) {
	obj, src := firstObject(t, `{ properties: 5 }`)
	if got := jsonSchemaToParams(obj, src); got != nil {
		t.Fatalf("expected nil when properties not object, got %v", got)
	}
}
