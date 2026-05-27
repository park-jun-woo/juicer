//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 1회 수렴 패스: 모든 파일의 include_router prefix를 전파하고 변경 여부를 반환한다
package fastapi

// propagatePrefixPass runs one convergence pass over all files, propagating
// include_router prefixes. It returns true if any source file's prefix changed.
func propagatePrefixPass(absRoot string, files []fileInfo,
	fileByPath map[string]*fileInfo, origSnapshot map[string]map[string]string) bool {

	changed := false
	for i := range files {
		if propagateFileIncludes(absRoot, &files[i], fileByPath, origSnapshot) {
			changed = true
		}
	}
	return changed
}
