//ff:func feature=ddl type=test control=sequence
//ff:what applyCreateTable 빈 라인/컬럼명 없는 라인 스킵 분기 테스트
package ddl

import "testing"

func TestApplyCreateTableSkipBlankCol(t *testing.T) {
	tables := map[string]*Table{}
	// trailing comma yields an empty line; a comment-only fragment yields no col name
	stmt := `CREATE TABLE Foo (
  id integer,
  ,
  name text
)`
	applyCreateTable(tables, "Foo", stmt)
	tbl, ok := tables["foo"]
	if !ok {
		t.Fatalf("table not registered: %v", tables)
	}
	if len(tbl.Columns) != 2 {
		t.Errorf("blank line must be skipped, got cols %+v", tbl.Columns)
	}
}
