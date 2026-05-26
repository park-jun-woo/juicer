//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what 부모 DTO 이름으로 import 경로를 추적하여 scanner.Field 목록을 반환한다
package nestjs

import (
	"path/filepath"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// resolveParentDTOFields resolves a parent DTO by name and returns scanner.Field slice.
func resolveParentDTOFields(parentName, referrerFile string, imports map[string]string, projectRoot string, cache map[string][]scanner.Field) []scanner.Field {
	if cached, ok := cache[parentName]; ok {
		return cached
	}
	importPath, ok := imports[parentName]
	if !ok {
		return nil
	}
	referrerDir := filepath.Dir(referrerFile)
	absPath := resolveImportPath(referrerDir, importPath, projectRoot)
	if absPath == "" {
		return nil
	}
	fields, err := extractDTO(absPath, parentName, imports, projectRoot, cache)
	if err != nil {
		return nil
	}
	return fields
}
