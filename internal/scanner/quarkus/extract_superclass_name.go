//ff:func feature=scan type=extract control=iteration dimension=1 topic=quarkus
//ff:what superclass 노드에서 부모 클래스명을 추출한다
package quarkus

import sitter "github.com/smacker/go-tree-sitter"

func extractSuperclassName(superclass *sitter.Node, src []byte) string {
	for i := 0; i < int(superclass.ChildCount()); i++ {
		child := superclass.Child(i)
		if child.Type() == "type_identifier" {
			return nodeText(child, src)
		}
	}
	return ""
}
