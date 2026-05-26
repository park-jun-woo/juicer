//ff:func feature=ddl type=test control=sequence
//ff:what ADD CONSTRAINT FK 시 같은 컬럼의 이름 없는 테이블 레벨 FK가 제거되는지 통합 테스트
package ddl

import (
	"strings"
	"testing"
)

func TestApplyAddConstraint_RemovesUnnamedFK(t *testing.T) {
	tbl := &Table{
		Name: "contract_tenants",
		Columns: []Column{
			{Name: "id", Raw: "id BIGSERIAL PRIMARY KEY"},
			{Name: "tenant_id", Raw: "tenant_id BIGINT NOT NULL"},
		},
		Constraints: []string{
			"FOREIGN KEY (tenant_id) REFERENCES tenants(id)",
		},
	}

	applyAddConstraint(tbl, "CONSTRAINT contract_tenants_tenant_id_fkey FOREIGN KEY (tenant_id) REFERENCES users(id)")

	// Only the new named FK should remain
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d: %v", len(tbl.Constraints), tbl.Constraints)
	}
	if !strings.Contains(tbl.Constraints[0], "CONSTRAINT contract_tenants_tenant_id_fkey") {
		t.Fatalf("expected named constraint, got %q", tbl.Constraints[0])
	}
	if !strings.Contains(tbl.Constraints[0], "REFERENCES users(id)") {
		t.Fatalf("expected new FK target, got %q", tbl.Constraints[0])
	}
}
