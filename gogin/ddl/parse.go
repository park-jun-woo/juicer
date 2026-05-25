//ff:func feature=ddl type=parse control=iteration dimension=1
//ff:what *.up.sql 파일을 순차 적용하여 최종 테이블 상태 반환
package ddl

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// Parse reads all *.up.sql files in dir, applies them sequentially,
// and returns the final table state.
func Parse(dir string) (map[string]*Table, error) {
	files, err := filepath.Glob(filepath.Join(dir, "*.up.sql"))
	if err != nil {
		return nil, fmt.Errorf("glob: %w", err)
	}
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
