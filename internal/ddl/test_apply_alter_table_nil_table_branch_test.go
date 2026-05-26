//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_NilTableBranch 테스트
package ddl

import "testing"

func TestApplyAlterTable_NilTableBranch(t *testing.T) {
	tables := map[string]*Table{}
	applyAlterTable(tables, "nonexistent", "ADD COLUMN email TEXT")
}
