//ff:func feature=ddl type=test control=sequence
//ff:what TestHasColumn_NotFoundBranch 테스트
package ddl

import "testing"

func TestHasColumn_NotFoundBranch(t *testing.T) {
	tbl := &Table{Columns: []Column{{Name: "id"}}}
	if hasColumn(tbl, "email") {
		t.Fatal("expected false")
	}
}
