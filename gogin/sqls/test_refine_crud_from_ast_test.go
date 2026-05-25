//ff:func feature=sql type=parse control=sequence
//ff:what TestRefineCRUDFromAST 테스트
package sqls

import (
	"testing"
)

func TestRefineCRUDFromAST(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		got := refineCRUDFromAST(nil)
		if got != "EXEC" {
			t.Errorf("expected EXEC, got %q", got)
		}
	})

	t.Run("with INSERT inline", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Create(ctx context.Context) { r.db.ExecContext(ctx, "INSERT INTO users (name) VALUES ($1)") }
`
		body := parseMethodBody(t, src, "Create")
		got := refineCRUDFromAST(body)
		if got != "INSERT" {
			t.Errorf("expected INSERT, got %q", got)
		}
	})

	t.Run("with UPDATE inline", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Update(ctx context.Context) { r.db.ExecContext(ctx, "UPDATE users SET name = $1") }
`
		body := parseMethodBody(t, src, "Update")
		got := refineCRUDFromAST(body)
		if got != "UPDATE" {
			t.Errorf("expected UPDATE, got %q", got)
		}
	})

	t.Run("with DELETE inline", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Delete(ctx context.Context) { r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1") }
`
		body := parseMethodBody(t, src, "Delete")
		got := refineCRUDFromAST(body)
		if got != "DELETE" {
			t.Errorf("expected DELETE, got %q", got)
		}
	})

	t.Run("no SQL in ExecContext", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Do(ctx context.Context) { r.db.ExecContext(ctx, "not sql") }
`
		body := parseMethodBody(t, src, "Do")
		got := refineCRUDFromAST(body)
		if got != "EXEC" {
			t.Errorf("expected EXEC, got %q", got)
		}
	})

	t.Run("non-ExecContext call", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Do(ctx context.Context) { r.db.QueryContext(ctx, "SELECT 1") }
`
		body := parseMethodBody(t, src, "Do")
		got := refineCRUDFromAST(body)
		if got != "EXEC" {
			t.Errorf("expected EXEC for non-ExecContext, got %q", got)
		}
	})
}
