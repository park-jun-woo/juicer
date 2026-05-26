//ff:func feature=scan type=test control=sequence
//ff:what TestFieldToProperty_FormatCov 테스트
package scanner

import "testing"

func TestFieldToProperty_FormatCov(t *testing.T) {
	f := Field{Name: "CreatedAt", Type: "time.Time"}
	prop := fieldToProperty(f)
	if prop["format"] == nil {
		t.Fatal("expected format hint for time.Time")
	}
}
