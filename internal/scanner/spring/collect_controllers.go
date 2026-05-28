//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 파싱된 Java 파일에서 @RestController 클래스를 수집한다
package spring

func collectControllers(files []*fileInfo) []controllerInfo {
	var result []controllerInfo
	for _, fi := range files {
		controllers := extractControllers(fi)
		result = append(result, controllers...)
	}
	return result
}
