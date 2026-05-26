//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_NestedStructCov 테스트
package scanner

import "testing"

func TestFieldToProperty_NestedStructCov(t *testing.T) {
	f := Field{Name: "Address", Type: "Address", Fields: []Field{{Name: "city", JSON: "city", Type: "string"}}}
	prop := fieldToProperty(f)
	if prop["type"] != "object" {
		t.Fatalf("expected object, got %v", prop["type"])
	}
}
