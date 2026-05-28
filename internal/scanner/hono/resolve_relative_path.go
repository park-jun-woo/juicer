//ff:func feature=scan type=extract control=iteration dimension=1 topic=hono
//ff:what 상대 경로 import를 실제 파일 경로로 해석한다
package hono

import (
	"os"
	"path/filepath"
)

func resolveRelativePath(dir, importPath string) string {
	abs := filepath.Join(dir, importPath)
	if _, err := os.Stat(abs); err == nil {
		return abs
	}
	extensions := []string{".ts", ".tsx", "/index.ts", "/index.tsx"}
	for _, ext := range extensions {
		candidate := abs + ext
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	return ""
}
