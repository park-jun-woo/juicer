//ff:func feature=ddl type=extract control=sequence
//ff:what TestHasColumn_Empty 테스트
package ddl

import "testing"

func TestHasColumn_Empty(t *testing.T) {
	tbl := &Table{}
	if hasColumn(tbl, "id") {
		t.Fatal("expected false")
	}
}
