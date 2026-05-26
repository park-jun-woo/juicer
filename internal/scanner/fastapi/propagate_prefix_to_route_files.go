//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what include_router 체인을 통해 원본 라우트 파일에 축적된 prefix를 전파한다
package fastapi

// propagatePrefixToRouteFiles handles the case where a router defined in one
// file (e.g., views.py) is included in another file (e.g., __init__.py) with a
// prefix. The route decorators are in views.py but the prefix is only known
// in __init__.py. This function traces include_router chains and propagates the
// accumulated prefix back to the file where routes are defined.
func propagatePrefixToRouteFiles(absRoot string, files []fileInfo, globalPrefixes map[string]map[string]string) {
	fileByPath := make(map[string]*fileInfo, len(files))
	for i := range files {
		fileByPath[files[i].absPath] = &files[i]
	}

	for i := range files {
		fi := &files[i]
		includes := findIncludeRouterCalls(fi.root, fi.src)
		if len(includes) == 0 {
			continue
		}
		importMap := buildRouterImportMap(absRoot, fi)
		for _, inc := range includes {
			propagateSingleInclude(fi, inc, importMap, fileByPath)
		}
	}
}
