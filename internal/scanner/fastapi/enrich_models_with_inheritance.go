//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what cross-file 상속 2차 패스로 부모가 모델인 클래스를 추가 등록한다
package fastapi

// enrichModelsWithInheritance adds classes whose parent is already a known model.
// This handles cross-file inheritance like LoginUser(BaseSchema) where BaseSchema
// is defined in another file and already collected in globalModels.
func enrichModelsWithInheritance(files []fileInfo, globalModels map[string]*fileInfo) {
	for tryInheritAllFiles(files, globalModels) {
	}
}
