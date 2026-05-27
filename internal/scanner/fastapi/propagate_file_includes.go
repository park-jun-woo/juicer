//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 단일 파일의 include_router 호출들을 순회하며 prefix를 전파한다
package fastapi

// propagateFileIncludes iterates over the include_router calls found in fi and
// propagates the accumulated prefix to each source file. It returns true if any
// source file's prefix was changed.
func propagateFileIncludes(absRoot string, fi *fileInfo,
	fileByPath map[string]*fileInfo, origSnapshot map[string]map[string]string) bool {

	includes := findIncludeRouterCalls(fi.root, fi.src)
	if len(includes) == 0 {
		return false
	}
	importMap := buildRouterImportMap(absRoot, fi, includes)
	changed := false
	for _, inc := range includes {
		if propagateSingleInclude(fi, inc, importMap, fileByPath, origSnapshot) {
			changed = true
		}
	}
	return changed
}
