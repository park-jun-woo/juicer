//ff:func feature=ddl type=test control=sequence
//ff:what applyAlterClause에서 멀티라인 ADD CONSTRAINT FK가 본문 포함하여 정상 저장되는지 테스트
package ddl

import "testing"

func TestApplyAlterClause_MultilineAddConstraintFK(t *testing.T) {
	tbl := &Table{Name: "contract_tenants"}

	clause := "ADD CONSTRAINT contract_tenants_tenant_id_fkey\n  FOREIGN KEY (tenant_id)\n  REFERENCES tenants (id)"
	applyAlterClause(tbl, clause)
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	want := "CONSTRAINT contract_tenants_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES tenants (id)"
	if tbl.Constraints[0] != want {
		t.Fatalf("constraint = %q, want %q", tbl.Constraints[0], want)
	}
}
