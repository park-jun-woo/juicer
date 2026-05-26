//ff:func feature=ddl type=test control=sequence
//ff:what TestApplyStatement_CreateTable 테스트
package ddl

import "testing"

func TestApplyStatement_CreateTable(t *testing.T) {
	tables := make(map[string]*Table)
	applyStatement(tables, "CREATE TABLE users (id INT PRIMARY KEY)")
	if tables["users"] == nil {
		t.Fatal("expected users table")
	}
}
