//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldToProperty_Pointer 테스트
package scanner

import "testing"

func TestFieldToProperty_Pointer(t *testing.T) {
	f := Field{Name: "Age", Type: "*int"}
	prop := fieldToProperty(f)
	if prop["type"] != "integer" {
		t.Fatalf("expected integer, got %v", prop["type"])
	}
}
