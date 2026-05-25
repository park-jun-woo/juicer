//ff:func feature=ddl type=parse control=sequence
//ff:what TestApplyCreateTableConstraints CREATE TABLE 제약 조건 및 edge case 테스트
package ddl

import (
	"testing"
)

func TestApplyCreateTableConstraints(t *testing.T) {
	t.Run("table with PRIMARY KEY constraint", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE t (id BIGINT, PRIMARY KEY (id))"
		applyCreateTable(tables, "t", stmt)

		tbl := tables["t"]
		if len(tbl.Columns) != 1 {
			t.Errorf("expected 1 column, got %d", len(tbl.Columns))
		}
		if len(tbl.Constraints) != 1 {
			t.Errorf("expected 1 constraint, got %d", len(tbl.Constraints))
		}
	})

	t.Run("table with UNIQUE constraint", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE t (id BIGINT, email TEXT, UNIQUE (email))"
		applyCreateTable(tables, "t", stmt)

		tbl := tables["t"]
		if len(tbl.Constraints) != 1 {
			t.Errorf("expected 1 constraint, got %d", len(tbl.Constraints))
		}
	})

	t.Run("table with CHECK constraint", func(t *testing.T) {
		tables := make(map[string]*Table)
		stmt := "CREATE TABLE t (id BIGINT, age INT, CHECK (age > 0))"
		applyCreateTable(tables, "t", stmt)

		tbl := tables["t"]
		if len(tbl.Constraints) != 1 {
			t.Errorf("expected 1 constraint, got %d", len(tbl.Constraints))
		}
	})

	t.Run("empty lines in body", func(t *testing.T) {
		tables := make(map[string]*Table)
		// Body with trailing comma produces empty element after split
		stmt := "CREATE TABLE t (\n    id BIGINT,\n    \n)"
		applyCreateTable(tables, "t", stmt)

		tbl := tables["t"]
		if len(tbl.Columns) != 1 {
			t.Errorf("expected 1 column, got %d", len(tbl.Columns))
		}
	})

	t.Run("comment-only column line", func(t *testing.T) {
		tables := make(map[string]*Table)
		// A line that is only a comment - extractColumnName returns ""
		stmt := "CREATE TABLE t (\n    -- just a comment\n    id BIGINT\n)"
		applyCreateTable(tables, "t", stmt)

		tbl := tables["t"]
		// The comment-only line should be skipped
		if len(tbl.Columns) < 1 {
			t.Errorf("expected at least 1 column, got %d", len(tbl.Columns))
		}
	})
}
