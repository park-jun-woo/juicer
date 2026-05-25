//ff:func feature=scan type=extract control=sequence
//ff:what TestFieldToProperty_String 테스트
package scanner

import "testing"

func TestFieldToProperty_String(t *testing.T) {
	f := Field{Name: "Name", Type: "string"}
	prop := fieldToProperty(f)
	if prop["type"] != "string" {
		t.Fatalf("expected string, got %v", prop["type"])
	}
}
