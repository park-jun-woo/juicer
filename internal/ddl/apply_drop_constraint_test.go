//ff:func feature=ddl type=test control=sequence
//ff:what applyDropConstraint의 이름 기반 Constraints 제거 테스트
package ddl

import "testing"

func TestApplyDropConstraint(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Constraints: []string{
			"CONSTRAINT users_email_key UNIQUE (email)",
			"CONSTRAINT users_pkey PRIMARY KEY (id)",
		},
	}

	applyDropConstraint(tbl, "users_email_key")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint, got %d", len(tbl.Constraints))
	}
	if tbl.Constraints[0] != "CONSTRAINT users_pkey PRIMARY KEY (id)" {
		t.Fatalf("wrong constraint remaining: %q", tbl.Constraints[0])
	}

	// Drop non-existent (no-op)
	applyDropConstraint(tbl, "nonexistent")
	if len(tbl.Constraints) != 1 {
		t.Fatalf("expected 1 constraint after no-op, got %d", len(tbl.Constraints))
	}
}
