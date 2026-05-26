//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import 정보에서 타입명→소스파일 맵을 생성한다
package fastapi

// buildImportMap creates a map from type name to source file path.
func buildImportMap(imports []importInfo, referrerDir string) map[string]string {
	result := make(map[string]string)
	for _, imp := range imports {
		resolved := resolveImportPath(referrerDir, imp.module)
		if resolved != "" {
			result[imp.name] = resolved
		}
	}
	return result
}
