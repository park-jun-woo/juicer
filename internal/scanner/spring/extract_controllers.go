//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what fileInfo에서 RestController 클래스를 추출한다
package spring

func extractControllers(fi *fileInfo) []controllerInfo {
	var result []controllerInfo
	classes := findAllByType(fi.root, "class_declaration")
	for _, cls := range classes {
		if !isRestController(cls, fi.src) {
			continue
		}
		ci := buildControllerInfo(cls, fi)
		result = append(result, ci)
	}
	return result
}
