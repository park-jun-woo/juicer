//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_RenameNilBranch 테스트
package ddl

import "testing"

func TestApplyAlterTable_RenameNilBranch(t *testing.T) {
	tables := map[string]*Table{}
	applyAlterTable(tables, "nonexistent", "RENAME TO people")
}
