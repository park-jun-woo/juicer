//ff:func feature=scan type=extract control=iteration dimension=1 topic=django
//ff:what choices 키워드 인자에서 선택값 목록을 추출한다
package django

import sitter "github.com/smacker/go-tree-sitter"

// extractChoiceValues extracts choice values from a choices keyword argument.
func extractChoiceValues(kwNode *sitter.Node, src []byte) []string {
	listNode := findChildByType(kwNode, "list")
	if listNode == nil {
		return nil
	}
	var choices []string
	for j := 0; j < int(listNode.ChildCount()); j++ {
		c := listNode.Child(j)
		if c.Type() == "string" {
			choices = append(choices, unquotePython(nodeText(c, src)))
		}
	}
	if len(choices) == 0 {
		return nil
	}
	return choices
}
