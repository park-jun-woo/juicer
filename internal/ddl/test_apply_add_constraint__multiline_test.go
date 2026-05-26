//ff:func feature=ddl type=test control=sequence
//ff:what 멀티라인 UNIQUE CONSTRAINT 정의가 단일 행으로 정규화되어 저장되는지 테스트
package ddl

import "testing"

func TestApplyAddConstraint_MultilineUnique(t *testing.T) {
	tbl := &Table{Name: "tenant_billings"}

	applyAddConstraint(tbl, "CONSTRAINT uq_billing_contract_month_type\n  UNIQUE (contract_id, billing_month, charge_type)")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	want := "CONSTRAINT uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)"
	if tbl.Constraints[0] != want {
		t.Fatalf("constraint = %q, want %q", tbl.Constraints[0], want)
	}
}
