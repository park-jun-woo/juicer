//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_PointerCov 테스트
package scanner

import "testing"

func TestFieldToProperty_PointerCov(t *testing.T) {
	f := Field{Name: "Count", Type: "*int"}
	prop := fieldToProperty(f)
	if prop["type"] != "integer" {
		t.Fatalf("expected integer, got %v", prop["type"])
	}
}
