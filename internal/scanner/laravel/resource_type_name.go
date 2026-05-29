//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what object_creation 노드가 Resource/Collection 타입이면 그 이름을 반환한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func resourceTypeName(oc *sitter.Node, src []byte) string {
	nameNode := findChildByType(oc, "name")
	if nameNode == nil {
		return ""
	}
	resName := nodeText(nameNode, src)
	if strings.HasSuffix(resName, "Resource") || strings.HasSuffix(resName, "Collection") {
		return resName
	}
	return ""
}
