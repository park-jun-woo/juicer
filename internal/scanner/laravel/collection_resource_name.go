//ff:func feature=scan type=extract control=sequence topic=laravel
//ff:what scoped 호출이 Resource::collection() 패턴이면 리소스 이름을 반환한다
package laravel

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func collectionResourceName(sc *sitter.Node, src []byte) string {
	nameNode := findChildByType(sc, "name")
	if nameNode == nil {
		return ""
	}
	name := nodeText(nameNode, src)
	if secondScopedName(sc, src) != "collection" {
		return ""
	}
	if strings.HasSuffix(name, "Resource") || strings.HasSuffix(name, "Collection") {
		return name
	}
	return ""
}
