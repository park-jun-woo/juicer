//ff:func feature=sql type=parse control=sequence
//ff:what TestCollectInlineSQLArgs 테스트
package sqls

import (
	"testing"
)

func TestCollectInlineSQLArgs(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		got := collectInlineSQLArgs(nil)
		if got != nil {
			t.Error("expected nil")
		}
	})

	t.Run("with inline SQL in ExecContext", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Do(ctx context.Context) { r.db.ExecContext(ctx, "INSERT INTO users (name) VALUES ($1)") }
`
		body := parseMethodBody(t, src, "Do")
		got := collectInlineSQLArgs(body)
		if len(got) != 1 {
			t.Errorf("expected 1 fragment, got %d: %v", len(got), got)
		}
	})

	t.Run("non-DB call skipped", func(t *testing.T) {
		src := `package test
import "fmt"

type R struct{}
func (r *R) Do() { fmt.Println("SELECT id FROM users WHERE name = $1") }
`
		body := parseMethodBody(t, src, "Do")
		got := collectInlineSQLArgs(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments for non-DB call, got %d", len(got))
		}
	})

	t.Run("short inline string skipped", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Do(ctx context.Context) { r.db.ExecContext(ctx, "short") }
`
		body := parseMethodBody(t, src, "Do")
		got := collectInlineSQLArgs(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments for short string, got %d", len(got))
		}
	})

	t.Run("non-SQL inline string skipped", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Do(ctx context.Context) { r.db.ExecContext(ctx, "this is not sql but long enough string") }
`
		body := parseMethodBody(t, src, "Do")
		got := collectInlineSQLArgs(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments for non-SQL string, got %d", len(got))
		}
	})

	t.Run("non-selector call skipped", func(t *testing.T) {
		src := `package test

func doSomething(s string) {}
func f() { doSomething("SELECT id FROM users WHERE name = $1") }
`
		body := parseMethodBody(t, src, "f")
		got := collectInlineSQLArgs(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments, got %d", len(got))
		}
	})
}
