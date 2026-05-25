//ff:func feature=sql type=parse control=sequence
//ff:what TestCollectSQLFragments 테스트
package sqls

import (
	"testing"
)

func TestCollectSQLFragments(t *testing.T) {
	t.Run("nil body", func(t *testing.T) {
		got := collectSQLFragments(nil)
		if got != nil {
			t.Error("expected nil")
		}
	})

	t.Run("with backtick SQL", func(t *testing.T) {
		src := "package test\nfunc f() {\n\tq := `SELECT id, name FROM users WHERE id = $1`\n\t_ = q\n}\n"
		body := parseMethodBody(t, src, "f")
		got := collectSQLFragments(body)
		if len(got) != 1 {
			t.Errorf("expected 1 fragment, got %d: %v", len(got), got)
		}
	})

	t.Run("short backtick string skipped", func(t *testing.T) {
		src := "package test\nfunc f() {\n\tq := `short`\n\t_ = q\n}\n"
		body := parseMethodBody(t, src, "f")
		got := collectSQLFragments(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments for short string, got %d", len(got))
		}
	})

	t.Run("no SQL keywords", func(t *testing.T) {
		src := "package test\nfunc f() {\n\tq := `this is a long enough string without any sql keywords at all`\n\t_ = q\n}\n"
		body := parseMethodBody(t, src, "f")
		got := collectSQLFragments(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments for non-SQL string, got %d", len(got))
		}
	})

	t.Run("double-quoted string skipped", func(t *testing.T) {
		src := "package test\nfunc f() {\n\tq := \"SELECT id FROM users WHERE name = $1\"\n\t_ = q\n}\n"
		body := parseMethodBody(t, src, "f")
		got := collectSQLFragments(body)
		if len(got) != 0 {
			t.Errorf("expected 0 fragments for double-quoted string, got %d", len(got))
		}
	})
}
