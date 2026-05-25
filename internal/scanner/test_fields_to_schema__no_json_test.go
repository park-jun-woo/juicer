//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldsToSchema_NoJSON 테스트
package scanner

import "testing"

func TestFieldsToSchema_NoJSON(t *testing.T) {
	fields := []Field{{Name: "ID", Type: "int"}}
	schema := fieldsToSchema(fields)
	props := schema["properties"].(map[string]any)
	if _, ok := props["ID"]; !ok {
		t.Fatal("expected ID in properties")
	}
}
