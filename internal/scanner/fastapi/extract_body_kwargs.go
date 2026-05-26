//ff:func feature=scan type=extract control=sequence topic=fastapi
//ff:what Body(...) 호출 노드에서 alias, embed 키워드 인자를 추출한다
package fastapi

import sitter "github.com/smacker/go-tree-sitter"

// extractBodyKwargs extracts alias and embed keyword arguments from a Body(...) call
// node found within the given parameter node.
func extractBodyKwargs(param *sitter.Node, src []byte) (alias string, embed bool) {
	callNode := findCallNode(param)
	if callNode == nil {
		return
	}
	args := findChildByType(callNode, "argument_list")
	if args == nil {
		return
	}
	alias = extractKeywordArg(args, "alias", src)
	embedVal := extractKeywordArg(args, "embed", src)
	if embedVal == "True" {
		embed = true
	}
	return
}
