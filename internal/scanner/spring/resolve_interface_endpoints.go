//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 인터페이스 파일에서 HTTP 엔드포인트를 추출한다
package spring

func resolveInterfaceEndpoints(ifacePath, ifaceName, projectRoot string) (string, []endpointInfo) {
	fi, err := parseFile(projectRoot, ifacePath)
	if err != nil {
		return "", nil
	}
	ifaces := findAllByType(fi.root, "interface_declaration")
	for _, iface := range ifaces {
		nameNode := findChildByType(iface, "identifier")
		if nameNode == nil || nodeText(nameNode, fi.src) != ifaceName {
			continue
		}
		prefix := extractClassPrefix(iface, fi.src)
		endpoints := extractInterfaceMethodEndpoints(iface, fi)
		return prefix, endpoints
	}
	return "", nil
}
