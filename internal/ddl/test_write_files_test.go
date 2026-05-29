//ff:func feature=ddl type=parse control=sequence
//ff:what TestWriteFiles 테스트
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles(t *testing.T) {
	t.Run("writes files", func(t *testing.T) {
		dir := t.TempDir()
		tables := map[string]*Table{
			"users": {
				Name:    "users",
				Columns: []Column{{Name: "id", Raw: "id BIGINT PRIMARY KEY"}},
			},
			"posts": {
				Name:    "posts",
				Columns: []Column{{Name: "id", Raw: "id BIGINT"}},
			},
		}
		if err := WriteFiles(nil, tables, dir); err != nil {
			t.Fatalf("WriteFiles() error: %v", err)
		}

		// Check files exist (numeric prefix added by WriteFiles)
		for _, suffix := range []string{"*_users.sql", "*_posts.sql"} {
			matches, _ := filepath.Glob(filepath.Join(dir, suffix))
			if len(matches) == 0 {
				t.Errorf("expected a %s file in %s", suffix, dir)
				continue
			}
			data, err := os.ReadFile(matches[0])
			if err != nil {
				t.Errorf("ReadFile(%s) error: %v", matches[0], err)
				continue
			}
			if len(data) == 0 {
				t.Errorf("file %s is empty", matches[0])
			}
		}
	})

	t.Run("creates output directory", func(t *testing.T) {
		base := t.TempDir()
		dir := filepath.Join(base, "subdir", "output")
		tables := map[string]*Table{
			"t": {Name: "t", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		}
		if err := WriteFiles(nil, tables, dir); err != nil {
			t.Fatalf("WriteFiles() error: %v", err)
		}
		matches, _ := filepath.Glob(filepath.Join(dir, "*_t.sql"))
		if len(matches) == 0 {
			t.Errorf("expected a *_t.sql file to exist in %s", dir)
		}
	})

	t.Run("mkdir error", func(t *testing.T) {
		// Use a path that cannot be created
		err := WriteFiles(nil, map[string]*Table{
			"t": {Name: "t"},
		}, "/dev/null/impossible")
		if err == nil {
			t.Error("expected error for impossible directory")
		}
	})

	t.Run("write error", func(t *testing.T) {
		dir := t.TempDir()
		outDir := filepath.Join(dir, "out")
		// Create the output directory first (so MkdirAll succeeds)
		os.MkdirAll(outDir, 0o755)
		// Now make it read-only so WriteFile fails
		os.Chmod(outDir, 0o555)
		t.Cleanup(func() { os.Chmod(outDir, 0o755) })

		err := WriteFiles(nil, map[string]*Table{
			"t": {Name: "t", Columns: []Column{{Name: "id", Raw: "id INT"}}},
		}, outDir)
		if err == nil {
			t.Error("expected error for read-only directory")
		}
	})
}
