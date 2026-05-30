//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what jsonSchemaPropertyToField 테스트
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

func TestJSONSchemaPropertyToField_NonObject(t *testing.T) {
	// value is a string literal, not an object -> only name/json set
	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	f := jsonSchemaPropertyToField("flag", n, src)
	if f.Name != "flag" || f.JSON != "flag" {
		t.Errorf("name/json wrong: %+v", f)
	}
	if f.Type != "" {
		t.Errorf("expected empty type for non-object, got %q", f.Type)
	}
}

func TestJSONSchemaPropertyToField_Array(t *testing.T) {
	obj, src := firstObject(t, `{ type: "array", items: { type: "string" } }`)
	f := jsonSchemaPropertyToField("tags", obj, src)
	if f.Type != "string[]" {
		t.Errorf("expected string[], got %q", f.Type)
	}
}
