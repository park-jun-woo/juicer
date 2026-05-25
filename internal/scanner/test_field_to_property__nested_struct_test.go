//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldToProperty_NestedStruct 테스트
package scanner

import "testing"

func TestFieldToProperty_NestedStruct(t *testing.T) {
	f := Field{Name: "Addr", Type: "Address", Fields: []Field{{Name: "City", Type: "string"}}}
	prop := fieldToProperty(f)
	if prop["type"] != "object" {
		t.Fatalf("expected object, got %v", prop["type"])
	}
}
