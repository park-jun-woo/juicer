//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what import 문에서 타입명의 소스 파일 경로를 추적한다
package nestjs

import (
	"os"
	"path/filepath"
	"strings"
)

// resolveImportPath resolves a relative import path to an absolute .ts file path.
// referrerDir is the directory of the file containing the import statement.
func resolveImportPath(referrerDir, importPath string) string {
	if !strings.HasPrefix(importPath, ".") {
		return "" // non-relative imports not supported
	}

	base := filepath.Join(referrerDir, importPath)

	// Try with .ts extension
	if !strings.HasSuffix(base, ".ts") {
		candidate := base + ".ts"
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
		// Try index.ts
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
