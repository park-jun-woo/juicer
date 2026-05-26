//ff:func feature=ddl type=test control=sequence
//ff:what TestHasColumn_NotFoundCov 테스트
package ddl

import "testing"

func TestHasColumn_NotFoundCov(t *testing.T) {
	tbl := &Table{Columns: []Column{{Name: "id"}}}
	if hasColumn(tbl, "missing") {
		t.Fatal("expected false")
	}
}
