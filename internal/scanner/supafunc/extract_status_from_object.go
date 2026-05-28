//ff:func feature=scan type=extract control=iteration dimension=1 topic=supafunc
//ff:what object 노드의 pair에서 status 프로퍼티의 number 값을 추출한다
package supafunc

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func extractStatusFromObject(obj *sitter.Node, src []byte) string {
	pairs := childrenOfType(obj, "pair")
	for _, pair := range pairs {
		key := findChildByType(pair, "property_identifier")
		if key == nil || nodeText(key, src) != "status" {
			continue
		}
		val := findChildByType(pair, "number")
		if val == nil {
			continue
		}
		return strings.TrimSpace(nodeText(val, src))
	}
	return ""
}
