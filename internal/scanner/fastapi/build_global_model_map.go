//ff:func feature=scan type=extract control=iteration dimension=1 topic=fastapi
//ff:what 모든 파일에서 모델명→fileInfo 맵을 생성한다
package fastapi

// buildGlobalModelMap builds a map from model name to the fileInfo that defines it.
func buildGlobalModelMap(files []fileInfo) map[string]*fileInfo {
	globalModels := make(map[string]*fileInfo)
	for i := range files {
		for name := range files[i].models {
			globalModels[name] = &files[i]
		}
	}
	return globalModels
}
