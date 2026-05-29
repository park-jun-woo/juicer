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

		// Check files exist
		for _, name := range []string{"users.sql", "posts.sql"} {
			path := filepath.Join(dir, name)
			data, err := os.ReadFile(path)
			if err != nil {
				t.Errorf("ReadFile(%s) error: %v", name, err)
				continue
			}
			if len(data) == 0 {
				t.Errorf("file %s is empty", name)
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
		if _, err := os.Stat(filepath.Join(dir, "t.sql")); err != nil {
			t.Errorf("expected t.sql to exist: %v", err)
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
