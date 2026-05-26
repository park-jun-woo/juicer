//ff:func feature=scan type=extract control=sequence topic=nestjs
//ff:what DTO 타입명을 소스 파일에서 추적하여 필드를 해석한다
package nestjs

import (
	"path/filepath"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// resolveDTOFields resolves a DTO type to its fields.
func resolveDTOFields(dr dtoRequest, cache map[string][]scanner.Field) ([]scanner.Field, error) {
	if cached, ok := cache[dr.typeName]; ok {
		return cached, nil
	}
	importPath, ok := dr.imports[dr.typeName]
	if !ok {
		return nil, nil
	}
	referrerDir := filepath.Dir(dr.referrer)
	absPath := resolveImportPath(referrerDir, importPath, dr.projectRoot)
	if absPath == "" {
		return nil, nil
	}
	fields, err := extractDTO(absPath, dr.typeName, dr.imports, dr.projectRoot, cache)
	if err != nil {
		return nil, err
	}
	cache[dr.typeName] = fields
	return fields, nil
}
