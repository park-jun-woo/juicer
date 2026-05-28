//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 컨트롤러의 인터페이스 목록에서 엔드포인트를 추출한다
package spring

func resolveControllerInterfaceEndpoints(ci *controllerInfo, fi *fileInfo) {
	for _, ifaceName := range ci.interfaces {
		ifacePath := findInterfaceFile(ifaceName, fi.imports, fi.absPath, fi.projectRoot)
		if ifacePath == "" {
			continue
		}
		ifacePrefix, eps := resolveInterfaceEndpoints(ifacePath, ifaceName, fi.projectRoot)
		if len(eps) == 0 {
			continue
		}
		if ci.prefix == "" && ifacePrefix != "" {
			ci.prefix = ifacePrefix
		}
		ci.endpoints = append(ci.endpoints, eps...)
	}
}
