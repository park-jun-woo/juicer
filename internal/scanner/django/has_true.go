//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what keyword_argument 노드에 True 값이 있는지 확인한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// hasTrue checks if a keyword_argument node contains a True value.
func hasTrue(kwNode *sitter.Node, src []byte) bool {
	for j := 0; j < int(kwNode.ChildCount()); j++ {
		v := kwNode.Child(j)
		if nodeText(v, src) == "True" {
			return true
		}
	}
	return false
}
