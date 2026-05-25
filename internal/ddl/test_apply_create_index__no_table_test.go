//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateIndex_NoTable 테스트
package ddl

import "testing"

func TestApplyCreateIndex_NoTable(t *testing.T) {
	tables := map[string]*Table{}
	applyCreateIndex(tables, "nonexistent", "CREATE INDEX idx ON nonexistent (x)")
}
