//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractFieldNameAndType 테스트
package spring

import "testing"

func TestExtractFieldNameAndType(t *testing.T) {
	field, src := firstField(t, `class D { private String name; }`)
	if extractFieldName(field, src) != "name" || extractFieldType(field, src) != "String" {
		t.Fatal("field name/type")
	}
}
