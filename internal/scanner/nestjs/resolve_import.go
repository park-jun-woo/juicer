//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what import 문에서 타입명의 소스 파일 경로를 추적한다
package nestjs

import (
	"path/filepath"
	"strings"
)

// resolveImportPath resolves a relative import path to an absolute .ts file path.
// referrerDir is the directory of the file containing the import statement.
// projectRoot is used for non-relative project-internal imports (e.g. 'src/users/dto/create-user.dto').
func resolveImportPath(referrerDir, importPath string, projectRoot ...string) string {
	if strings.HasPrefix(importPath, ".") {
		return tryResolveTS(filepath.Join(referrerDir, importPath))
	}
	if len(projectRoot) > 0 && projectRoot[0] != "" {
		return resolveNonRelativeImport(projectRoot[0], importPath)
	}
	return ""
}
