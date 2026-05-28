//ff:func feature=scan type=extract control=iteration dimension=1 topic=spring
//ff:what 노드의 modifiers에 지정한 키워드가 모두 포함되어 있는지 확인한다
package spring

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

func hasModifiers(node *sitter.Node, src []byte, required ...string) bool {
	mods := findChildByType(node, "modifiers")
	if mods == nil {
		return false
	}
	text := nodeText(mods, src)
	for _, req := range required {
		if !strings.Contains(text, req) {
			return false
		}
	}
	return true
}
