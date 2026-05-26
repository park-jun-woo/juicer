//ff:func feature=ddl type=test control=sequence
//ff:what removeUnnamedFK가 이름 있는 FK는 유지하는지 테스트
package ddl

import "testing"

func TestRemoveUnnamedFK_KeepsNamedFK(t *testing.T) {
	tbl := &Table{
		Name: "contract_tenants",
		Constraints: []string{
			"CONSTRAINT contract_tenants_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES users(id)",
		},
	}

	removeUnnamedFK(tbl, "tenant_id")

	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
}
