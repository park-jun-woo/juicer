//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldToProperty_Array 테스트
package scanner

import "testing"

func TestFieldToProperty_Array(t *testing.T) {
	f := Field{Name: "Tags", Type: "[]string"}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}
}
