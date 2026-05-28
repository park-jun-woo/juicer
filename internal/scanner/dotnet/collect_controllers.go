//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 파싱된 C# 파일에서 [ApiController] 클래스를 수집한다
package dotnet

func collectControllers(files []*fileInfo) []controllerInfo {
	var result []controllerInfo
	for _, fi := range files {
		controllers := extractControllers(fi)
		result = append(result, controllers...)
	}
	return result
}
