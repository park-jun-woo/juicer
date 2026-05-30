//ff:func feature=scan type=test control=iteration dimension=1 topic=actix
//ff:what findSerdeTokenTree 테스트 헬퍼: attr 아이템 목록에서 첫 serde token_tree 노드 조회
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

// findSerdeTokenTree returns the first serde token_tree node among items.
func findSerdeTokenTree(items []*sitter.Node, src []byte) *sitter.Node {
	for _, it := range items {
		if tt := serdeTokenTree(it, src); tt != nil {
			return tt
		}
	}
	return nil
}
