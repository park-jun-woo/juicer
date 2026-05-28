//ff:func feature=scan type=extract control=iteration dimension=1 topic=dotnet
//ff:what 어트리뷰트 리스트에서 특정 이름의 어트리뷰트를 찾는다
package dotnet

import sitter "github.com/smacker/go-tree-sitter"

func findAttributeInList(attrList *sitter.Node, src []byte, name string) *sitter.Node {
	for _, attr := range childrenOfType(attrList, "attribute") {
		if attributeName(attr, src) == name {
			return attr
		}
	}
	return nil
}
