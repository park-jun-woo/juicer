//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 확장자 없는 경로에 .ts 또는 /index.ts를 붙여 실제 파일 경로를 반환한다
package express

import (
	"os"
	"path/filepath"
	"strings"
)

func resolveExtension(base string) string {
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
