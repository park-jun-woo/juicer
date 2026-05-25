//ff:func feature=sql type=parse control=sequence
//ff:what TestDetectCRUD 테스트
package sqls

import (
	"testing"
)

func TestDetectCRUD(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		got := detectCRUD(nil)
		if got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("QueryContext", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Get(ctx context.Context) { r.db.QueryContext(ctx, "SELECT 1") }
`
		body := parseMethodBody(t, src, "Get")
		got := detectCRUD(body)
		if got != "SELECT" {
			t.Errorf("expected SELECT, got %q", got)
		}
	})

	t.Run("non-selector call in body", func(t *testing.T) {
		src := `package test
import "fmt"

type R struct{}
func (r *R) Do() { fmt.Println("hello") }
`
		body := parseMethodBody(t, src, "Do")
		got := detectCRUD(body)
		if got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})

	t.Run("ExecContext", func(t *testing.T) {
		src := `package test
import "context"
import "database/sql"

type R struct{ db *sql.DB }
func (r *R) Do(ctx context.Context) { r.db.ExecContext(ctx, "DELETE FROM users") }
`
		body := parseMethodBody(t, src, "Do")
		got := detectCRUD(body)
		if got != "EXEC" {
			t.Errorf("expected EXEC, got %q", got)
		}
	})
}
