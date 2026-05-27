//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what include_router 체인을 통해 원본 라우트 파일에 축적된 prefix를 전파한다
package fastapi

// propagatePrefixToRouteFiles handles the case where a router defined in one
// file (e.g., views.py) is included in another file (e.g., __init__.py) with a
// prefix. The route decorators are in views.py but the prefix is only known
// in __init__.py. This function traces include_router chains and propagates the
// accumulated prefix back to the file where routes are defined.
func propagatePrefixToRouteFiles(absRoot string, files []fileInfo) {
	fileByPath := make(map[string]*fileInfo, len(files))
	for i := range files {
		fileByPath[files[i].absPath] = &files[i]
	}

	// 전파 시작 전 스냅샷: merge + resolveDotted 완료 후의 정확한 prefix
	// globalPrefixes와 달리 이 시점의 fi.prefixes는 dotted resolution 완료 상태
	origSnapshot := buildGlobalPrefixMap(files)

	for pass := 0; pass < 10; pass++ {
		if !propagatePrefixPass(absRoot, files, fileByPath, origSnapshot) {
			break
		}
	}
}
