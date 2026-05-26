//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what import된 라우터 변수의 prefix를 글로벌 맵에서 해석하여 로컬 맵에 병합한다
package fastapi

// mergeImportedRouterPrefixes iterates over all files and for each
// include_router call referencing an imported router variable, merges the
// cross-file prefix into the local prefix map.
func mergeImportedRouterPrefixes(absRoot string, files []fileInfo, globalPrefixes map[string]map[string]string) {
	for i := range files {
		fi := &files[i]
		includes := findIncludeRouterCalls(fi.root, fi.src)
		if len(includes) == 0 {
			continue
		}
		importMap := buildRouterImportMap(absRoot, fi)
		for _, inc := range includes {
			mergeSingleInclude(fi, inc, importMap, globalPrefixes)
		}
	}
}
