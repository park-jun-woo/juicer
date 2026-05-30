//ff:func feature=ddl type=test control=sequence
//ff:what TestStripColumnReferences_NoMatch 테스트
package ddl

import "testing"

func TestStripColumnReferences_NoMatch(t *testing.T) {
	tbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "id", Raw: "id BIGINT REFERENCES owners(id)"},
		},
	}

	stripColumnReferences(tbl, "missing")
	if tbl.Columns[0].Raw != "id BIGINT REFERENCES owners(id)" {
		t.Fatalf("expected no change, got %q", tbl.Columns[0].Raw)
	}
}
