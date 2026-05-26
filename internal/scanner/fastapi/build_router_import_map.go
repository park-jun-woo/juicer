//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import 정보에서 로컬 이름 → 소스 파일 경로 맵을 생성한다
package fastapi

import "path/filepath"

// buildRouterImportMap builds a map from local imported name to the absPath of
// the file that defines it. Both relative and absolute imports are resolved.
func buildRouterImportMap(absRoot string, fi *fileInfo) map[string]string {
	importMap := make(map[string]string)
	referrerDir := filepath.Dir(fi.absPath)
	for _, imp := range fi.imports {
		resolved := resolveImportPath(referrerDir, imp.module)
		if resolved == "" {
			resolved = resolveAbsoluteImportPath(absRoot, imp.module)
		}
		if resolved != "" {
			importMap[imp.name] = resolved
		}
	}
	return importMap
}
