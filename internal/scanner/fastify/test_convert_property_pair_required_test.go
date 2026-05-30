//ff:func feature=scan type=test control=iteration dimension=1 topic=fastify
//ff:what TestConvertPropertyPair_Required 테스트
package fastify

import "testing"

func TestConvertPropertyPair_Required(t *testing.T) {
	pairs, src := schemaPairs(t, `{ name: { type: "string" }, age: { type: "integer" } }`)
	requiredSet := map[string]bool{"name": true}

	validateByName := map[string]string{}
	for _, p := range pairs {
		f := convertPropertyPair(p, src, requiredSet)
		if f == nil {
			t.Fatal("nil field")
		}
		validateByName[f.Name] = f.Validate
	}
	if v, ok := validateByName["name"]; !ok || v != "required" {
		t.Errorf("name should be required, got %q (present=%v)", v, ok)
	}
	if v, ok := validateByName["age"]; !ok || v == "required" {
		t.Errorf("age should not be required, got %q (present=%v)", v, ok)
	}
}
