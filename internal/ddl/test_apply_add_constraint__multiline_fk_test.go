//ff:func feature=ddl type=test control=sequence
//ff:what 멀티라인 FOREIGN KEY CONSTRAINT 정의가 단일 행으로 정규화되어 저장되는지 테스트
package ddl

import "testing"

func TestApplyAddConstraint_MultilineForeignKey(t *testing.T) {
	tbl := &Table{Name: "building_owners"}

	applyAddConstraint(tbl, "CONSTRAINT building_owners_owner_id_fkey\n  FOREIGN KEY (owner_id)\n  REFERENCES owners (id)")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	want := "CONSTRAINT building_owners_owner_id_fkey FOREIGN KEY (owner_id) REFERENCES owners (id)"
	if tbl.Constraints[0] != want {
		t.Fatalf("constraint = %q, want %q", tbl.Constraints[0], want)
	}
}
