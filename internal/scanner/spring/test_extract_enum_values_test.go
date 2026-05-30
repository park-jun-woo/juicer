//ff:func feature=scan type=test control=sequence topic=spring
//ff:what TestExtractEnumValues 테스트
package spring

import "testing"

func TestExtractEnumValues(t *testing.T) {
	root, src := parseS(t, `enum Status { OPEN, CLOSED }`)
	en := findAllByType(root, "enum_declaration")[0]
	vals := extractEnumValues(en, src)
	if len(vals) != 2 || vals[0] != "OPEN" {
		t.Fatalf("got %v", vals)
	}
}
