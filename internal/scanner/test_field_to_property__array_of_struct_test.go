//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldToProperty_ArrayOfStruct 테스트
package scanner

import "testing"

func TestFieldToProperty_ArrayOfStruct(t *testing.T) {
	f := Field{Name: "Items", Type: "[]Item", Fields: []Field{{Name: "ID", Type: "int"}}}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}
}
