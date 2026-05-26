//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_NilTableCov 테스트
package ddl

import "testing"

func TestApplyAlterTable_NilTableCov(t *testing.T) {
	tables := map[string]*Table{}
	applyAlterTable(tables, "missing", "ADD COLUMN foo TEXT")
}
