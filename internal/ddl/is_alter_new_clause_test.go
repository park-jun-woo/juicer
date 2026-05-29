//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what TestIsAlterNewClause_AddColumn 테스트
package ddl

import "testing"

func TestIsAlterNewClause_AddColumn(t *testing.T) {
	// every recognised clause prefix must return true (case-insensitive)
	truthy := []string{
		"ADD COLUMN name TEXT",
		"DROP COLUMN name",
		"ADD CONSTRAINT pk PRIMARY KEY (id)",
		"DROP CONSTRAINT pk",
		"ALTER COLUMN name SET NOT NULL",
		"RENAME TO users2",
		"add column lower_case TEXT",
	}
	for _, c := range truthy {
		if !isAlterNewClause(c) {
			t.Fatalf("expected true for %q", c)
		}
	}

	// non-clause continuation text must return false
	if isAlterNewClause("REFERENCES users(id)") {
		t.Fatal("expected false for non-clause text")
	}
}
