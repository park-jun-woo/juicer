//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyCreateIndex_NilTableCov 테스트
package ddl

import "testing"

func TestApplyCreateIndex_NilTableCov(t *testing.T) {
	tables := map[string]*Table{}
	applyCreateIndex(tables, "missing", "CREATE INDEX idx ON missing (col)")
}
