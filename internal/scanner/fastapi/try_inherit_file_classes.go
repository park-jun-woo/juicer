//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 1개 파일의 클래스를 순회하며 상속 기반 모델 등록을 시도한다
package fastapi

// tryInheritFileClasses iterates over classes in fi and registers any
// whose parent is already a known model. Returns true if any class was added.
func tryInheritFileClasses(fi *fileInfo, globalModels map[string]*fileInfo) bool {
	if fi.root == nil {
		return false
	}
	added := false
	classes := findAllByType(fi.root, "class_definition")
	for _, cls := range classes {
		if tryInheritClass(cls, fi, globalModels) {
			added = true
		}
	}
	return added
}
