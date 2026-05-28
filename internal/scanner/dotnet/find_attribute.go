//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 노드에서 특정 이름의 어트리뷰트를 찾아 반환한다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func findAttribute(node *sitter.Node, src []byte, name string) *sitter.Node {
	for _, attrList := range childrenOfType(node, "attribute_list") {
		attr := findAttributeInList(attrList, src, name)
		if attr != nil {
			return attr
		}
	}
	return nil
}
