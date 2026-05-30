//ff:func feature=scan type=test control=sequence topic=quarkus
//ff:what TestExtractEnumValues 테스트
package quarkus

import "testing"

func TestExtractEnumValues(t *testing.T) {
	fi := qFileInfo(t, `enum Status { OPEN, CLOSED, PENDING }`)
	cls := findAllByType(fi.root, "enum_declaration")[0]
	vals := extractEnumValues(cls, fi.src)
	if len(vals) != 3 || vals[0] != "OPEN" {
		t.Fatalf("got %v", vals)
	}
}
