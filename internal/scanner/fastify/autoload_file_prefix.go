//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastify
//ff:what autoload 디렉터리 구조에서 파일의 prefix를 계산한다 (서브디렉터리명=세그먼트, 비-index 파일명=세그먼트)
package fastify

import (
	"path/filepath"
	"strings"
)

func autoloadFilePrefix(baseDir, path, basePrefix string) (string, bool) {
	name := filepath.Base(path)
	if !strings.HasSuffix(name, ".ts") || strings.HasSuffix(name, ".d.ts") {
		return "", false
	}
	rel, err := filepath.Rel(baseDir, path)
	if err != nil {
		return "", false
	}
	prefix := basePrefix
	for _, seg := range autoloadSegments(rel) {
		prefix = joinFastifyPath(prefix, seg)
	}
	return prefix, true
}
