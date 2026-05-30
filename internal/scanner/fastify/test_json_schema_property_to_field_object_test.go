//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaPropertyToField_Object 테스트
package fastify

import "testing"

func TestJSONSchemaPropertyToField_Object(t *testing.T) {
	obj, src := firstObject(t, `{ type: "string", format: "email" }`)
	f := jsonSchemaPropertyToField("email", obj, src)
	if f.Name != "email" || f.JSON != "email" {
		t.Errorf("name/json wrong: %+v", f)
	}
	if f.Type != "email" {
		t.Errorf("expected format-applied type email, got %q", f.Type)
	}
}
