//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what @action 데코레이터의 키워드 인자를 actionInfo에 적용한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// applyActionKeywords applies keyword arguments to actionInfo.
func applyActionKeywords(ai *actionInfo, args *sitter.Node, src []byte) {
	for i := 0; i < int(args.ChildCount()); i++ {
		child := args.Child(i)
		if child.Type() != "keyword_argument" {
			continue
		}
		keyNode := findChildByType(child, "identifier")
		if keyNode == nil {
			continue
		}
		applyOneActionKeyword(ai, nodeText(keyNode, src), child, src)
	}
}
