//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractFieldNameAndType 테스트
package quarkus

import "testing"

func TestExtractFieldNameAndType(t *testing.T) {
	root, src := parseQ(t, `class D { private String name; }`)
	fields := findAllByType(root, "field_declaration")
	if got := extractFieldName(fields[0], src); got != "name" {
		t.Fatalf("name: %q", got)
	}
	if got := extractFieldType(fields[0], src); got != "String" {
		t.Fatalf("type: %q", got)
	}
}
