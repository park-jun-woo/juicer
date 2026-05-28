//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 같은 파일 내부에 정의된 인터페이스를 해석한다
package spring

import "os"

func resolveSameFileInterface(referrerPath, ifaceName string) string {
	src, err := os.ReadFile(referrerPath)
	if err != nil {
		return ""
	}
	root, err := parseJava(src)
	if err != nil {
		return ""
	}
	ifaces := findAllByType(root, "interface_declaration")
	for _, iface := range ifaces {
		nameNode := findChildByType(iface, "identifier")
		if nameNode != nil && nodeText(nameNode, src) == ifaceName {
			return referrerPath
		}
	}
	return ""
}
