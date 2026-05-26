//ff:func feature=ddl type=test control=sequence
//ff:what removeUnnamedFK 단위 테스트
package ddl

import "testing"

func TestRemoveUnnamedFK(t *testing.T) {
	tbl := &Table{
		Name: "contract_tenants",
		Constraints: []string{
			"FOREIGN KEY (tenant_id) REFERENCES tenants(id)",
		},
	}

	removeUnnamedFK(tbl, "tenant_id")

	if len(tbl.Constraints) != 0 {
		t.Fatalf("expected 0 constraints, got %d: %v", len(tbl.Constraints), tbl.Constraints)
	}
}
