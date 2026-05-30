//ff:func feature=scan type=test control=sequence topic=fastify
//ff:what TestJSONSchemaPropertyToField_NonObject 테스트
package fastify

import "testing"

func TestJSONSchemaPropertyToField_NonObject(t *testing.T) {

	n, src := firstNodeOfType(t, `const x = "lit";`+"\n", "string")
	f := jsonSchemaPropertyToField("flag", n, src)
	if f.Name != "flag" || f.JSON != "flag" {
		t.Errorf("name/json wrong: %+v", f)
	}
	if f.Type != "" {
		t.Errorf("expected empty type for non-object, got %q", f.Type)
	}
}
