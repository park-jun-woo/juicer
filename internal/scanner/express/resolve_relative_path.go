//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 상대 경로를 절대 .ts 파일 경로로 해석한다
package express

import (
	"os"
	"path/filepath"
	"strings"
)

func resolveRelativePath(dir, importPath string) string {
	if !strings.HasPrefix(importPath, ".") {
		return ""
	}
	base := filepath.Join(dir, importPath)
	if !strings.HasSuffix(base, ".ts") {
		candidate := base + ".ts"
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		candidate = filepath.Join(base, "index.ts")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}
	if _, err := os.Stat(base); err == nil {
		return base
	}
	return ""
}
