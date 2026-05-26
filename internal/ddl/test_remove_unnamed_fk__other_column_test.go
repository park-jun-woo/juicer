//ff:func feature=ddl type=test control=sequence
//ff:what removeUnnamedFK가 다른 컬럼의 이름 없는 FK는 유지하는지 테스트
package ddl

import "testing"

func TestRemoveUnnamedFK_KeepsOtherColumn(t *testing.T) {
	tbl := &Table{
		Name: "contract_tenants",
		Constraints: []string{
			"FOREIGN KEY (contract_id) REFERENCES contracts(id)",
			"FOREIGN KEY (tenant_id) REFERENCES tenants(id)",
		},
	}

	removeUnnamedFK(tbl, "tenant_id")

	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d: %v", len(tbl.Constraints), tbl.Constraints)
	}
	if tbl.Constraints[0] != "FOREIGN KEY (contract_id) REFERENCES contracts(id)" {
		t.Fatalf("wrong constraint preserved: %q", tbl.Constraints[0])
	}
}
