//ff:func feature=scan type=extract control=sequence topic=express
//ff:what 상대 경로를 절대 TS/JS 소스 파일 경로로 해석한다 (확장자 후보 .ts/.tsx/.js/.jsx/.mjs/.cjs + index.<ext>)
package express

import (
	"path/filepath"
	"strings"
)

func resolveRelativePath(dir, importPath string) string {
	if !strings.HasPrefix(importPath, ".") {
		return ""
	}
	base := filepath.Join(dir, importPath)
	return resolveSourceBase(base)
}
