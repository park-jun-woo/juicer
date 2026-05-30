//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumAssignmentValue_NumberValueUsesName 테스트
package nestjs

import "testing"

func TestExtractEnumAssignmentValue_NumberValueUsesName(t *testing.T) {
	src := []byte(`enum E { FIRST = 1 }`)
	root, _ := parseTypeScript(src)
	asn := findAllByType(root, "enum_assignment")
	if len(asn) == 0 {
		t.Skip("no enum_assignment")
	}
	v, ok := extractEnumAssignmentValue(asn[0], src)
	if !ok || v != "FIRST" {
		t.Fatalf("got %q %v", v, ok)
	}
}
