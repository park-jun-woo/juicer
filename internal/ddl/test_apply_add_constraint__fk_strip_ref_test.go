//ff:func feature=ddl type=test control=sequence
//ff:what ADD CONSTRAINT FK 시 Column.Raw의 이전 REFERENCES가 제거되는지 통합 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestApplyAddConstraint_FKStripsInlineRef(t *testing.T) {
	tbl := &Table{
		Name: "building_owners",
		Columns: []Column{
			{Name: "id", Raw: "id BIGSERIAL PRIMARY KEY"},
			{Name: "owner_id", Raw: "owner_id BIGINT NOT NULL REFERENCES owners(id)"},
		},
	}

	// Simulate ADD CONSTRAINT with new FK target
	applyAddConstraint(tbl, "CONSTRAINT building_owners_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES users(id)")

	// Column.Raw must no longer contain the old REFERENCES
	wantRaw := "owner_id BIGINT NOT NULL"
	gotRaw := tbl.Columns[1].Raw
	if gotRaw != wantRaw {
		t.Fatalf("Column.Raw = %q, want %q", gotRaw, wantRaw)
	}

	// Constraint must be present with new FK
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	if !strings.Contains(tbl.Constraints[0], "REFERENCES users(id)") {
		t.Fatalf("constraint missing new FK: %q", tbl.Constraints[0])
	}
}
