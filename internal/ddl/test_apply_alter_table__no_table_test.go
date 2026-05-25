//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyAlterTable_NoTable 테스트
package ddl

import "testing"

func TestApplyAlterTable_NoTable(t *testing.T) {
	tables := map[string]*Table{}
	applyAlterTable(tables, "nonexistent", "ADD COLUMN x INT")
}
