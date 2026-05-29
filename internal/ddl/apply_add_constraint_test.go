//ff:func feature=ddl type=test control=sequence
//ff:what applyAddConstraint의 Constraints 추가 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestApplyAddConstraint(t *testing.T) {
	// Non-FK constraint: simply appended.
	tbl := &Table{Name: "tenant_billings"}
	applyAddConstraint(tbl, "uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	want := "uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)"
	if tbl.Constraints[0] != want {
		t.Fatalf("constraint = %q, want %q", tbl.Constraints[0], want)
	}

	// FK constraint: removes prior unnamed FK and strips inline REFERENCES.
	fkTbl := &Table{
		Name: "orders",
		Columns: []Column{
			{Name: "owner_id", Raw: "owner_id INT REFERENCES users(id)"},
		},
		Constraints: []string{
			"FOREIGN KEY (owner_id) REFERENCES users(id)",
		},
	}
	applyAddConstraint(fkTbl, "fk_orders_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE")

	// The unnamed FK must be gone, leaving only the new named FK.
	if len(fkTbl.Constraints) != 1 {
		t.Fatalf("expected unnamed FK removed leaving 1 constraint, got %d: %v", len(fkTbl.Constraints), fkTbl.Constraints)
	}
	if fkTbl.Constraints[0] != "fk_orders_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE" {
		t.Fatalf("unexpected constraint: %q", fkTbl.Constraints[0])
	}
	// Inline REFERENCES on the column must be stripped.
	if strings.Contains(fkTbl.Columns[0].Raw, "REFERENCES") {
		t.Fatalf("inline REFERENCES not stripped: %q", fkTbl.Columns[0].Raw)
	}
}
