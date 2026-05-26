//ff:func feature=scan type=test control=sequence
//ff:what TestFieldsToSchema_NoJSONCov 테스트
package scanner

import "testing"

func TestFieldsToSchema_NoJSONCov(t *testing.T) {
	fields := []Field{{Name: "Foo", Type: "string"}}
	schema := fieldsToSchema(fields)
	props := schema["properties"].(map[string]any)
	if props["Foo"] == nil {
		t.Fatal("expected Foo key (fallback to Name)")
	}
}
