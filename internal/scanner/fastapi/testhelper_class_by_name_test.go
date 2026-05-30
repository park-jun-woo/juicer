//ff:func feature=scan type=test control=iteration dimension=1 topic=fastapi
//ff:what classByName 테스트 헬퍼
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

func classByName(root *sitter.Node, src []byte, name string) *sitter.Node {
	for _, c := range findAllByType(root, "class_definition") {
		id := findChildByType(c, "identifier")
		if id != nil && nodeText(id, src) == name {
			return c
		}
	}
	return nil
}
