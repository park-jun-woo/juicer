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

func TestRemoveUnnamedFK_KeepsNonMatching(t *testing.T) {
	tbl := &Table{
		Name: "contract_tenants",
		Constraints: []string{
			// FK on a different column — should be kept.
			"FOREIGN KEY (other_id) REFERENCES others(id)",
			// Named FK on the target column — should be kept.
			"CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id)",
			// Unnamed FK on the target column — should be removed.
			"FOREIGN KEY (tenant_id) REFERENCES tenants(id)",
		},
	}

	removeUnnamedFK(tbl, "tenant_id")

	if len(tbl.Constraints) != 2 {
		t.Fatalf("expected 2 constraints, got %d: %v", len(tbl.Constraints), tbl.Constraints)
	}
	if tbl.Constraints[0] != "FOREIGN KEY (other_id) REFERENCES others(id)" {
		t.Errorf("expected non-matching FK to be kept, got %q", tbl.Constraints[0])
	}
	if tbl.Constraints[1] != "CONSTRAINT fk_tenant FOREIGN KEY (tenant_id) REFERENCES tenants(id)" {
		t.Errorf("expected named FK to be kept, got %q", tbl.Constraints[1])
	}
}
