//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 include_router 호출의 축적된 prefix를 원본 파일에 전파한다
package fastapi

// propagateSingleInclude propagates the accumulated prefix from an
// include_router call in fi to the source file where the router is defined.
// It returns true if the source file's prefix was changed.
func propagateSingleInclude(fi *fileInfo, inc includeCall, importMap map[string]string,
	fileByPath map[string]*fileInfo, origSnapshot map[string]map[string]string) bool {

	srcFile := importMap[inc.childVar]
	if srcFile == "" {
		return false
	}
	srcFI := fileByPath[srcFile]
	if srcFI == nil {
		return false
	}
	origVar := findOriginalVarName(inc.childVar, srcFI)
	if origVar == "" {
		return false
	}

	// 전파 직전 스냅샷에서 원래 로컬 값을 읽는다
	// (merge + resolveDotted 완료 후, 전파 시작 전의 정확한 값)
	origPrefixes := origSnapshot[fi.absPath]
	origLocalParent := origPrefixes[inc.parentVar]
	origChildAccum := origPrefixes[inc.childVar]

	// extra 기여분 추출: origChildAccum에서 origLocalParent를 제거
	extraContrib := stripLeadingPath(origChildAccum, origLocalParent)

	// 갱신된 부모 prefix + extra + 소스 파일의 원래 prefix
	parentPrefix := fi.prefixes[inc.parentVar]
	srcOrigPrefix := ""
	if sp := origSnapshot[srcFile]; sp != nil {
		srcOrigPrefix = sp[origVar]
	}
	accumulated := joinPath(parentPrefix, extraContrib, srcOrigPrefix)

	if accumulated == "" || srcFI.prefixes[origVar] == accumulated {
		return false
	}
	srcFI.prefixes[origVar] = accumulated
	return true
}
