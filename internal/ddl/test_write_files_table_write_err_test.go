//ff:func feature=ddl type=test control=iteration dimension=1
//ff:what WriteFiles enums+tables 정상 기록 및 테이블 파일 경로 검증
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFilesTableWriteErr(t *testing.T) {
	dir := t.TempDir()
	enums := []EnumType{{Name: "Role", Values: []string{"ADMIN"}}}
	tables := map[string]*Table{
		`"users"`: {Name: `"users"`, Columns: []Column{{Name: "id", Raw: "id integer"}}},
		"orgs":    {Name: "orgs", Columns: []Column{{Name: "id", Raw: "id integer"}}},
	}
	if err := WriteFiles(enums, tables, dir); err != nil {
		t.Fatalf("WriteFiles: %v", err)
	}
	entries, _ := os.ReadDir(dir)
	if len(entries) != 3 {
		t.Fatalf("want 3 files, got %d: %v", len(entries), entries)
	}
	// enum file first, table file name strips quotes
	if _, err := os.Stat(filepath.Join(dir, "00_enums.sql")); err != nil {
		t.Errorf("enum file: %v", err)
	}
	found := false
	for _, e := range entries {
		if e.Name() == "02_users.sql" || e.Name() == "01_users.sql" {
			found = true
		}
	}
	if !found {
		t.Errorf("quoted table file not sanitized: %v", entries)
	}
}
