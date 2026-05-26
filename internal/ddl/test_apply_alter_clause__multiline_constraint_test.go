//ff:func feature=ddl type=test control=sequence
//ff:what applyAlterClause에서 멀티라인 ADD CONSTRAINT UNIQUE가 본문 포함하여 정상 저장되는지 테스트
package ddl

import "testing"

func TestApplyAlterClause_MultilineAddConstraintUnique(t *testing.T) {
	tbl := &Table{Name: "tenant_billings"}

	clause := "ADD CONSTRAINT uq_billing_contract_month_type\n  UNIQUE (contract_id, billing_month, charge_type)"
	applyAlterClause(tbl, clause)
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	want := "CONSTRAINT uq_billing_contract_month_type UNIQUE (contract_id, billing_month, charge_type)"
	if tbl.Constraints[0] != want {
		t.Fatalf("constraint = %q, want %q", tbl.Constraints[0], want)
	}
}
