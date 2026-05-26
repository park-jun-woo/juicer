//ff:func feature=scan type=test control=sequence
//ff:what TestFieldsToSchema_RequiredCov 테스트
package scanner

import "testing"

func TestFieldsToSchema_RequiredCov(t *testing.T) {
	fields := []Field{{Name: "Email", JSON: "email", Type: "string", Validate: "required"}}
	schema := fieldsToSchema(fields)
	if schema["required"] == nil {
		t.Fatal("expected required field")
	}
}
