//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 모든 파일을 순회하며 상속 기반 모델 등록을 1회 시도한다
package fastapi

// tryInheritAllFiles runs one pass over all files, attempting to register
// classes whose parent is a known model. Returns true if any class was added.
func tryInheritAllFiles(files []fileInfo, globalModels map[string]*fileInfo) bool {
	added := false
	for i := range files {
		if tryInheritFileClasses(&files[i], globalModels) {
			added = true
		}
	}
	return added
}
