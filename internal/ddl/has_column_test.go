//ff:func feature=ddl type=test control=sequence
//ff:what TestHasColumn_Found 테스트
package ddl

import "testing"

func TestHasColumn_Found(t *testing.T) {
	tbl := &Table{Columns: []Column{{Name: "id"}, {Name: "name"}}}
	if !hasColumn(tbl, "id") {
		t.Fatal("expected true")
	}
}
