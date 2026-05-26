//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyCreateIndex_NilTable 테스트
package ddl

import "testing"

func TestApplyCreateIndex_NilTable(t *testing.T) {
	tables := map[string]*Table{}
	applyCreateIndex(tables, "nonexistent", "CREATE INDEX idx ON nonexistent (col)")
}
