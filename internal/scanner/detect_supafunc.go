//ff:func feature=scan type=extract control=iteration dimension=1
//ff:what supabase/functions/ 디렉토리와 index.ts 파일 존재 여부로 Supabase Edge Functions를 감지한다
package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func detectSupaFunc(root string) bool {
	funcDir := filepath.Join(root, "supabase", "functions")
	info, err := os.Stat(funcDir)
	if err != nil || !info.IsDir() {
		return false
	}
	entries, err := os.ReadDir(funcDir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), "_") {
			continue
		}
		indexPath := filepath.Join(funcDir, e.Name(), "index.ts")
		if _, err := os.Stat(indexPath); err == nil {
			return true
		}
	}
	return false
}
