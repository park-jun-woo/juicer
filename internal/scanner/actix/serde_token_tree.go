//ff:func feature=scan type=extract control=sequence topic=actix
//ff:what serde 어트리뷰트 항목에서 인자 token_tree 노드를 찾는다
package actix

import (
	sitter "github.com/smacker/go-tree-sitter"
)

func serdeTokenTree(attrItem *sitter.Node, src []byte) *sitter.Node {
	attr := findChildByType(attrItem, "attribute")
	if attr == nil {
		return nil
	}
	nameNode := findChildByType(attr, "identifier")
	if nameNode == nil {
		return nil
	}
	if nodeText(nameNode, src) != "serde" {
		return nil
	}
	return findChildByType(attr, "token_tree")
}
