//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import 정보에서 로컬 이름 → 소스 파일 경로 맵을 생성한다
package fastapi

import "path/filepath"

// buildRouterImportMap builds a map from local imported name to the absPath of
// the file that defines it. Both relative and absolute imports are resolved.
// When includes contain a childModule (e.g., items.router), the sub-module
// import path is resolved so that childVar maps to the correct source file.
func buildRouterImportMap(absRoot string, fi *fileInfo, includes []includeCall) map[string]string {
	importMap := make(map[string]string)
	referrerDir := filepath.Dir(fi.absPath)

	// 1. 기존: 단순 import name → 소스 파일 매핑
	for _, imp := range fi.imports {
		resolved := resolveImportPath(referrerDir, imp.module)
		if resolved == "" {
			resolved = resolveAbsoluteImportPath(absRoot, imp.module)
		}
		if resolved != "" {
			importMap[imp.name] = resolved
		}
	}

	// 2. childModule이 있는 include_router 호출 처리
	for _, inc := range includes {
		if inc.childModule == "" {
			continue
		}
		if _, exists := importMap[inc.childVar]; exists {
			continue
		}
		resolved := resolveChildModuleImport(absRoot, referrerDir, inc.childModule, fi.imports)
		if resolved != "" {
			importMap[inc.childVar] = resolved
		}
	}
	return importMap
}
