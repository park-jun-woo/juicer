//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_RenameNilTable 테스트
package ddl

import "testing"

func TestApplyAlterTable_RenameNilTable(t *testing.T) {
	tables := map[string]*Table{}
	applyAlterTable(tables, "nonexistent", "RENAME TO foo")
}
