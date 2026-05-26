//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_NestedArrayCov 테스트
package scanner

import "testing"

func TestFieldToProperty_NestedArrayCov(t *testing.T) {
	f := Field{Name: "Tags", Type: "[]Tag", Fields: []Field{{Name: "name", JSON: "name", Type: "string"}}}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}
}
