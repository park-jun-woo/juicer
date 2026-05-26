//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what 단일 include_router 호출의 축적된 prefix를 원본 파일에 전파한다
package fastapi

// propagateSingleInclude propagates the accumulated prefix from an
// include_router call in fi to the source file where the router is defined.
func propagateSingleInclude(fi *fileInfo, inc includeCall, importMap map[string]string, fileByPath map[string]*fileInfo) {
	srcFile := importMap[inc.childVar]
	if srcFile == "" {
		return
	}
	srcFI := fileByPath[srcFile]
	if srcFI == nil {
		return
	}
	accumulated := fi.prefixes[inc.childVar]
	if accumulated == "" {
		return
	}
	origVar := findOriginalVarName(inc.childVar, srcFI)
	if origVar == "" {
		return
	}
	if srcFI.prefixes[origVar] != accumulated {
		srcFI.prefixes[origVar] = accumulated
	}
}
