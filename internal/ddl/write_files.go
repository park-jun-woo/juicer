//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 테이블별 .sql 파일을 출력 디렉토리에 생성
package ddl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteFiles writes enum types and each table's DDL into outDir. Enum types are
// collected into a single file (sorted before table files) so dependent tables
// are applied after their types exist.
func WriteFiles(enums []EnumType, tables map[string]*Table, outDir string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", outDir, err)
	}

	if len(enums) > 0 {
		var sb strings.Builder
		for _, e := range sortedEnums(enums) {
			renderEnum(&sb, e)
		}
		path := filepath.Join(outDir, "0_enums.sql")
		if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}

	for _, t := range tables {
		var sb strings.Builder
		renderTable(&sb, t)
		path := filepath.Join(outDir, t.Name+".sql")
		if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}
