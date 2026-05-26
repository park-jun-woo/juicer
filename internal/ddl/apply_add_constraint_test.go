//ff:func feature=ddl type=test control=sequence
//ff:what applyAddConstraint의 Constraints 추가 테스트
package ddl

import "testing"

func TestApplyAddConstraint(t *testing.T) {
	tbl := &Table{Name: "tenant_billings"}

	applyAddConstraint(tbl, "uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	want := "uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)"
	if tbl.Constraints[0] != want {
		t.Fatalf("constraint = %q, want %q", tbl.Constraints[0], want)
	}
}
