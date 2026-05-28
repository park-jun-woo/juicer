//ff:func feature=scan type=extract control=sequence topic=django
//ff:what 대입문이 urlpatterns 대입인지 확인한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// isURLPatternsAssignment checks if an assignment node is `urlpatterns = [...]`.
func isURLPatternsAssignment(node *sitter.Node, src []byte) bool {
	leftNodes := childrenOfType(node, "identifier")
	if len(leftNodes) == 0 {
		return false
	}
	return nodeText(leftNodes[0], src) == "urlpatterns"
}
