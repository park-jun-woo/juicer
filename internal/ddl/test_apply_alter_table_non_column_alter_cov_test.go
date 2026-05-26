//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyAlterTable_NonColumnAlterCov 테스트
package ddl

import "testing"

func TestApplyAlterTable_NonColumnAlterCov(t *testing.T) {
	tables := map[string]*Table{
		"users": {Name: "users", Columns: []Column{{Name: "id", Raw: "id INT"}}},
	}
	applyAlterTable(tables, "users", "ADD CONSTRAINT pk PRIMARY KEY (id)")
}
