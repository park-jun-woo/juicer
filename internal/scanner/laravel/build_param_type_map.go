//ff:func feature=scan type=extract control=iteration dimension=1 topic=laravel
//ff:what 컨트롤러 메서드 파라미터들의 (이름 → OpenAPI 타입) 맵을 만든다
package laravel

func buildParamTypeMap(cm *controllerMethod) map[string]string {
	typeMap := make(map[string]string)
	for _, p := range cm.params {
		if p.typeName == "" || p.name == "request" {
			continue
		}
		oaType := phpTypeToOpenAPI(p.typeName)
		if oaType != "" {
			typeMap[p.name] = oaType
		}
	}
	return typeMap
}
