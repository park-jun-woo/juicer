//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what supabase/functions/ 하위 디렉토리에서 index.ts 파일을 수집한다 (_ 접두사 디렉토리 제외)
package supafunc

import (
	"os"
	"path/filepath"
	"strings"
)

func findFunctions(root string) ([]string, error) {
	funcDir := filepath.Join(root, "supabase", "functions")
	entries, err := os.ReadDir(funcDir)
	if err != nil {
		return nil, err
	}
	var files []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if strings.HasPrefix(e.Name(), "_") {
			continue
		}
		indexPath := filepath.Join(funcDir, e.Name(), "index.ts")
		if _, err := os.Stat(indexPath); err == nil {
			files = append(files, indexPath)
		}
	}
	return files, nil
}
