//ff:func feature=ddl type=extract control=sequence
//ff:what TestHasColumn_NotFound 테스트
package ddl

import "testing"

func TestHasColumn_NotFound(t *testing.T) {
	tbl := &Table{Columns: []Column{{Name: "id"}}}
	if hasColumn(tbl, "email") {
		t.Fatal("expected false")
	}
}
