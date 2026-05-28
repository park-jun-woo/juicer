//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what *.sql 파일을 순차 적용하여 최종 테이블 상태 반환 (*.down.sql 제외)
package ddl

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Parse reads all *.sql files in dir (excluding *.down.sql),
// applies them sequentially, and returns the final table state.
func Parse(dir string) (map[string]*Table, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.sql"))
	if err != nil {
		return nil, fmt.Errorf("glob: %w", err)
	}

	// Filter out down-migration files
	filtered := files[:0]
	for _, f := range files {
		base := filepath.Base(f)
		lower := strings.ToLower(base)
		if strings.HasSuffix(lower, ".down.sql") || strings.HasSuffix(lower, "_down.sql") {
			continue
		}
		filtered = append(filtered, f)
	}
	files = filtered

	sort.Strings(files)

	tables := make(map[string]*Table)

	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", f, err)
		}
		stmts := splitStatements(string(data))
		for _, stmt := range stmts {
			applyStatement(tables, stmt)
		}
	}

	return tables, nil
}
