//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_ArrayCov 테스트
package scanner

import "testing"

func TestFieldToProperty_ArrayCov(t *testing.T) {
	f := Field{Name: "IDs", Type: "[]int"}
	prop := fieldToProperty(f)
	if prop["type"] != "array" {
		t.Fatalf("expected array, got %v", prop["type"])
	}
}
