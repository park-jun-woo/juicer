//ff:func feature=ddl type=render control=iteration dimension=1
//ff:what 테이블별 .sql 파일을 출력 디렉토리에 생성
package ddl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteFiles writes each table's DDL to a separate .sql file in outDir.
func WriteFiles(tables map[string]*Table, outDir string) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return fmt.Errorf("mkdir %s: %w", outDir, err)
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
