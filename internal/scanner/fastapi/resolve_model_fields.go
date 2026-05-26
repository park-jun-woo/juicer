//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 모델 타입을 해석하여 필드를 반환한다
package fastapi

import (
	"path/filepath"

	"github.com/park-jun-woo/juicer/internal/scanner"
)

// resolveModelFields resolves a single model type to its fields.
func resolveModelFields(req modelRequest, cache map[string][]scanner.Field, globalModels map[string]*fileInfo) []scanner.Field {
	if cached, ok := cache[req.typeName]; ok {
		return cached
	}

	// Check global model map first (same or different file)
	if fi, ok := globalModels[req.typeName]; ok {
		if pyFields, ok := fi.models[req.typeName]; ok {
			fields := pydanticFieldsToScannerFields(pyFields)
			cache[req.typeName] = fields
			return fields
		}
	}

	// Try resolving via imports
	referrerDir := filepath.Dir(req.referrer)
	importMap := buildImportMap(req.imports, referrerDir)
	sourcePath, ok := importMap[req.typeName]
	if !ok {
		return nil
	}

	fields, err := extractPydanticModel(sourcePath, req.typeName)
	if err != nil || fields == nil {
		return nil
	}
	cache[req.typeName] = fields
	return fields
}
