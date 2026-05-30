//ff:func feature=ddl type=test control=sequence
//ff:what stripColumnReferences 단위 테스트
package ddl

import "testing"

func TestStripColumnReferences_Match(t *testing.T) {
	tbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "id", Raw: "id BIGINT"},
			{Name: "owner_id", Raw: "owner_id BIGINT NOT NULL REFERENCES owners(id) ON DELETE CASCADE"},
		},
	}
	stripColumnReferences(tbl, "owner_id")

	if tbl.Columns[1].Raw != "owner_id BIGINT NOT NULL" {
		t.Fatalf("expected inline references stripped, got %q", tbl.Columns[1].Raw)
	}
	// Untouched column stays the same.
	if tbl.Columns[0].Raw != "id BIGINT" {
		t.Fatalf("unrelated column changed: %q", tbl.Columns[0].Raw)
	}
}

func TestStripColumnReferences_CaseInsensitiveName(t *testing.T) {
	tbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "owner_id", Raw: "owner_id BIGINT REFERENCES owners(id)"},
		},
	}
	// colName supplied in upper case; lookup lowercases it.
	stripColumnReferences(tbl, "OWNER_ID")
	if tbl.Columns[0].Raw != "owner_id BIGINT" {
		t.Fatalf("expected stripped, got %q", tbl.Columns[0].Raw)
	}
}

func TestStripColumnReferences_NoMatch(t *testing.T) {
	tbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "id", Raw: "id BIGINT REFERENCES owners(id)"},
		},
	}
	// Column not present — loop completes without modifying anything.
	stripColumnReferences(tbl, "missing")
	if tbl.Columns[0].Raw != "id BIGINT REFERENCES owners(id)" {
		t.Fatalf("expected no change, got %q", tbl.Columns[0].Raw)
	}
}
