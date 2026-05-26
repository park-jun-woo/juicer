//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 기본 경로에서 .ts 파일 존재 여부를 확인하여 절대 경로를 반환한다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"
)

// tryResolveTS attempts to resolve a base path to an existing .ts file.
func tryResolveTS(base string) string {
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
