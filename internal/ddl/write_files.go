//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 테이블별 .sql 파일을 출력 디렉토리에 생성
package ddl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteFiles writes enum types and each table's DDL into outDir. File names get
// a fixed-width zero-padded numeric prefix so a lexicographic listing
// (e.g. `ls *.sql`) matches the FK-dependency apply order: enums first
// (00_enums.sql), then tables in topological order. The padding width is sized
// to the table count so all prefixes share the same width.
func WriteFiles(enums []EnumType, tables map[string]*Table, outDir string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", outDir, err)
	}

	width := prefixWidth(len(tables))

	if len(enums) > 0 {
		var sb strings.Builder
		for _, e := range sortedEnums(enums) {
			renderEnum(&sb, e)
		}
		name := fmt.Sprintf("%0*d_enums.sql", width, 0)
		path := filepath.Join(outDir, name)
		if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}

	for i, key := range topoSortTables(tables) {
		t := tables[key]
		var sb strings.Builder
		renderTable(&sb, t)
		// SQL content keeps quoted identifiers; only the file name is sanitized
		// so a quoted table name ("User") writes to NN_User.sql, not "User".sql.
		base := strings.ReplaceAll(t.Name, `"`, "")
		fileName := fmt.Sprintf("%0*d_%s.sql", width, i+1, base)
		path := filepath.Join(outDir, fileName)
		if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}
