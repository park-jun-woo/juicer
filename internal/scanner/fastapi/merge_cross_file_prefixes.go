//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 파일 간 include_router 체인을 추적하여 prefix를 병합한다
package fastapi

// mergeCrossFilePrefixes resolves cross-file include_router chains. For each
// file that contains include_router calls referencing imported router variables,
// this function traces the import back to the defining file, retrieves its
// prefix, and updates the local prefix map.
//
// After this pass, each fileInfo's prefixes map contains the fully resolved
// prefix chain including cross-file parents.
func mergeCrossFilePrefixes(absRoot string, files []fileInfo) {
	globalPrefixes := buildGlobalPrefixMap(files)
	mergeImportedRouterPrefixes(absRoot, files, globalPrefixes)
	propagatePrefixToRouteFiles(absRoot, files, globalPrefixes)
}
