//ff:func feature=scan type=extract control=sequence topic=fastify
//ff:what 상대 경로를 절대 .ts 파일 경로로 해석한다
package fastify

import (
	"path/filepath"
	"strings"
)

func resolveRelativePath(dir, importPath string) string {
	if !strings.HasPrefix(importPath, ".") {
		return ""
	}
	base := filepath.Join(dir, importPath)
	return resolveExtension(base)
}
