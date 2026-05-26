//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Python 상대 import 경로를 파일 시스템 경로로 변환한다
package fastapi

import (
	"os"
	"path/filepath"
	"strings"
)

// resolveImportPath resolves a Python import path to a file system path.
// referrerDir is the directory of the file containing the import.
// module is the import path (e.g., ".models" or "app.models").
func resolveImportPath(referrerDir, module string) string {
	if !strings.HasPrefix(module, ".") {
		return ""
	}

	dots := countLeadingDots(module)
	base := navigateUp(referrerDir, dots)

	modulePart := module[dots:]
	if modulePart == "" {
		return ""
	}

	relPath := strings.ReplaceAll(modulePart, ".", string(filepath.Separator))
	fullPath := filepath.Join(base, relPath)

	pyFile := fullPath + ".py"
	if _, err := os.Stat(pyFile); err == nil {
		return pyFile
	}
	initFile := filepath.Join(fullPath, "__init__.py")
	if _, err := os.Stat(initFile); err == nil {
		return initFile
	}
	return ""
}
