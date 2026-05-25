//ff:func feature=sql type=parse control=sequence
//ff:what TestParseFile Go 파일에서 SQL 메서드 파싱 테스트
package sqls

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile(t *testing.T) {
	dir := t.TempDir()

	t.Run("valid repo file", func(t *testing.T) {
		path := filepath.Join(dir, "user_repo.go")
		os.WriteFile(path, []byte(parseFileValidRepoSrc), 0o644)
		methods, err := parseFile(path)
		if err != nil {
			t.Fatalf("parseFile() error: %v", err)
		}
		if len(methods) < 1 {
			t.Errorf("expected at least 1 method, got %d", len(methods))
		}
	})

	t.Run("repo with no SQL methods", func(t *testing.T) {
		path := filepath.Join(dir, "empty_repo.go")
		os.WriteFile(path, []byte(parseFileEmptyRepoSrc), 0o644)
		methods, err := parseFile(path)
		if err != nil {
			t.Fatalf("parseFile() error: %v", err)
		}
		if len(methods) != 0 {
			t.Errorf("expected 0 methods, got %d", len(methods))
		}
	})

	t.Run("repo with non-method functions and vars", func(t *testing.T) {
		path := filepath.Join(dir, "misc_repo.go")
		os.WriteFile(path, []byte(parseFileMiscRepoSrc), 0o644)
		methods, err := parseFile(path)
		if err != nil {
			t.Fatalf("parseFile() error: %v", err)
		}
		if len(methods) != 0 {
			t.Errorf("expected 0 methods, got %d", len(methods))
		}
	})

	t.Run("repo with ExecContext INSERT", func(t *testing.T) {
		path := filepath.Join(dir, "order_repo.go")
		os.WriteFile(path, []byte(parseFileOrderRepoSrc), 0o644)
		methods, err := parseFile(path)
		if err != nil {
			t.Fatalf("parseFile() error: %v", err)
		}
		if len(methods) < 1 {
			t.Errorf("expected at least 1 method, got %d", len(methods))
		}
	})

	t.Run("repo with QueryRowContext and INSERT RETURNING", func(t *testing.T) {
		path := filepath.Join(dir, "item_repo.go")
		os.WriteFile(path, []byte(parseFileItemRepoSrc), 0o644)
		methods, err := parseFile(path)
		if err != nil {
			t.Fatalf("parseFile() error: %v", err)
		}
		for _, m := range methods {
			if m.Method == "CreateAndReturn" && m.CRUD != "INSERT" {
				t.Errorf("expected CRUD INSERT for CreateAndReturn, got %q", m.CRUD)
			}
		}
	})

	t.Run("invalid file", func(t *testing.T) {
		path := filepath.Join(dir, "bad.go")
		os.WriteFile(path, []byte("not go{{{"), 0o644)
		_, err := parseFile(path)
		if err == nil {
			t.Error("expected error for invalid Go file")
		}
	})
}
