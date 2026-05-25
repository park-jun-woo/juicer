//ff:func feature=sql type=parse control=iteration dimension=2
//ff:what queriesDir 내 .sql 파일에서 "-- name: queryName" 패턴 존재 여부 확인
package sqls

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// queryExists scans .sql files in queriesDir for a "-- name: queryName" pattern.
//
func queryExists(queriesDir, queryName string) bool {
	entries, err := os.ReadDir(queriesDir)
	if err != nil {
		return false
	}

	target := "-- name: " + queryName

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		path := filepath.Join(queriesDir, e.Name())
		f, err := os.Open(path)
		if err != nil {
			continue
		}
		sc := bufio.NewScanner(f)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())
			if strings.HasPrefix(line, target) {
				f.Close()
				return true
			}
		}
		f.Close()
	}
	return false
}

