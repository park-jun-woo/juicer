//ff:func feature=scan type=test control=sequence topic=nestjs
//ff:what TestExtractEnumAssignmentValue_PropertyIdentifier 테스트
package nestjs

import "testing"

func TestExtractEnumAssignmentValue_PropertyIdentifier(t *testing.T) {
	src := []byte(`enum E { Foo }`)
	root, _ := parseTypeScript(src)
	body := findAllByType(root, "enum_body")[0]

	pids := findAllByType(body, "property_identifier")
	if len(pids) == 0 {
		t.Skip("no property identifier")
	}
	v, ok := extractEnumAssignmentValue(pids[0], src)
	if !ok || v != "Foo" {
		t.Fatalf("got %q %v", v, ok)
	}
}
