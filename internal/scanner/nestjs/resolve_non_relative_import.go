//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 프로젝트 루트 기준 비상대 import 경로를 절대 경로로 변환한다
package nestjs

import (
	"path/filepath"
	"strings"
)

// resolveNonRelativeImport resolves project-internal non-relative imports.
// Handles patterns like 'src/users/dto/create-user.dto' and '@/decorators/field'.
func resolveNonRelativeImport(projectRoot, importPath string) string {
	path := importPath
	if strings.HasPrefix(path, "@/") {
		path = "src/" + path[2:]
	}
	base := filepath.Join(projectRoot, path)
	return tryResolveTS(base)
}
